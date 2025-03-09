package coroutine

import (
	"errors"
	"fmt"
	"gitee.com/quant1x/gox/logger"
	"gitee.com/quant1x/gox/runtime"
	"gitee.com/quant1x/gox/timestamp"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// 滑动窗口的锚点毫秒数, 默认0
	defaultRollingAnchorPoint = 0
	// 窗口相对于锚点的偏移量
	defaultRollingWindow = timestamp.MillisecondsPerDay
	// 相对于默认, 每天9点整
	defaultOffsetWindow = timestamp.MillisecondsPerHour * 9
)

var (
	ErrInvalidOffset = errors.New("offset must be in [0, 86400000)")
	ErrNilFunction   = errors.New("function cannot be nil")
)

// RollingOnce 实现基于时间窗口的周期性任务懒加载
// 默认每天在指定偏移时间（如9:00）重置状态，允许再次执行任务
type RollingOnce struct {
	name        string        // 任务名
	done        atomic.Bool   // 执行状态标志
	m           sync.Mutex    // 保护临界区
	initOnce    sync.Once     // 确保初始化只执行一次
	ticker      *time.Ticker  // 时间窗口检查定时器
	windowSize  atomic.Int64  // 窗口大小（毫秒）
	offset      atomic.Int64  // 相对于零点的偏移量（毫秒）
	observer    atomic.Int64  // 当前观察点时间戳（毫秒）
	task        func()        // 需要执行的任务函数
	closeSignal chan struct{} // 关闭信号通道
	closeOnce   sync.Once     // 确保关闭操作只执行一次
}

// NewRollingOnceWithHourAndMinute 创建指定时分的每日初始化RollingOnce实例
func NewRollingOnceWithHourAndMinute(hour, minute int, task func()) (*RollingOnce, error) {
	windowMs := defaultRollingWindow
	offsetMs := timestamp.MillisecondsPerHour*hour + timestamp.MillisecondsPerMinute*minute
	return NewRollingOnce(int64(windowMs), int64(offsetMs), task)
}

// NewRollingOnce 创建新的RollingOnce实例
// 参数说明:
//
//	windowMs - 窗口持续时间（毫秒），建议 >= 1000
//	offsetMs - 每日偏移量（毫秒），范围 [0, 86400000)
//	task     - 要执行的任务函数
func NewRollingOnce(windowMs, offsetMs int64, task func()) (*RollingOnce, error) {
	if windowMs < 1000 {
		return nil, fmt.Errorf("window size too small: %w", ErrInvalidOffset)
	}
	if offsetMs < 0 || offsetMs >= timestamp.MillisecondsPerDay {
		return nil, fmt.Errorf("invalid offset: %w", ErrInvalidOffset)
	}
	if task == nil {
		return nil, ErrNilFunction
	}

	ro := &RollingOnce{
		name:        runtime.FuncName(task),
		task:        task,
		closeSignal: make(chan struct{}, 1), // 缓冲通道防止阻塞
	}
	ro.windowSize.Store(windowMs)
	ro.offset.Store(offsetMs)
	ro.updateObserver()
	return ro, nil
}

// 更新观察点到下一个有效时间窗口
func (o *RollingOnce) updateObserver() {
	now := timestamp.Now()
	base := now - (now % timestamp.MillisecondsPerDay)
	newObserver := base + o.offset.Load()

	if now >= newObserver {
		newObserver += timestamp.MillisecondsPerDay
	}
	o.observer.Store(newObserver)
	logger.Debugf("[%s]窗口观察点更新为: %s", o.name, time.Unix(newObserver/1000, 0).Format(time.RFC3339)) // 修正此处
}

// Do 执行任务（仅在当前时间窗口内首次调用有效）
func (o *RollingOnce) Do() {
	// 双重检查锁定优化性能
	if o.done.Load() {
		return
	}

	o.initOnce.Do(func() {
		// 启动后台协程前先检查一次
		if o.windowExpired() {
			o.reset()
		}
		go o.watcher()
	})

	o.m.Lock()
	defer o.m.Unlock()

	if o.done.Load() {
		return
	}
	defer o.done.Store(true)

	o.task()
	logger.Infof("[%s]任务执行完成 @ %s", o.name, time.Now().Format(time.RFC3339))
}

// 判断当前窗口是否过期
func (o *RollingOnce) windowExpired() bool {
	return timestamp.Now() >= o.observer.Load()
}

// 后台监听协程
func (o *RollingOnce) watcher() {
	const checkInterval = 30 * time.Second // 平衡精度和性能
	o.ticker = time.NewTicker(checkInterval)
	defer o.ticker.Stop()

	for {
		select {
		case <-o.ticker.C:
			if o.windowExpired() {
				logger.Debugf("[%s]检测到窗口过期，触发重置", o.name)
				o.reset()
			}
		case <-o.closeSignal:
			logger.Debugf("[%s]接收到关闭信号，退出监听", o.name)
			return
		}
	}
}

// 重置执行状态
func (o *RollingOnce) reset() {
	o.m.Lock()
	defer o.m.Unlock()

	if timestamp.Now() < o.observer.Load() {
		return
	}

	o.done.Store(false)
	o.updateObserver()
	logger.Infof("[%s]状态已重置，下次执行窗口: %s", o.name, time.Unix(o.observer.Load()/1000, 0).Format(time.RFC3339)) // 修正此处
}

// Close 安全停止后台监听
func (o *RollingOnce) Close() {
	o.closeOnce.Do(func() {
		if o.ticker != nil {
			o.ticker.Stop()
		}
		close(o.closeSignal)
		logger.Info("资源已释放")
	})
}

// AdjustOffset 动态调整时间偏移量（线程安全）
func (o *RollingOnce) AdjustOffset(newOffsetMs int64) error {
	if newOffsetMs < 0 || newOffsetMs >= timestamp.MillisecondsPerDay {
		return ErrInvalidOffset
	}

	o.offset.Store(newOffsetMs)
	o.updateObserver()
	logger.Infof("[%s]时间偏移量已调整为: %dms", o.name, newOffsetMs)
	return nil
}
