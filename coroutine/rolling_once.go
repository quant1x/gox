package coroutine

import (
	"sync"
	"sync/atomic"
	"time"

	"gitee.com/quant1x/gox/logger"
	"gitee.com/quant1x/gox/runtime"
	"gitee.com/quant1x/gox/timestamp"
)

const (
	// 滑动窗口的锚点毫秒数, 默认0
	rollingAnchorPoint = 0
	// 窗口相对于锚点的偏移量
	rollingWindow = timestamp.MillisecondsPerDay
	// 相对于默认, 每天9点整
	offsetWindow = timestamp.MillisecondsPerHour * 9
)

// 计算下一个时间窗口
// 当前时间戳和当前观察点+偏移量比较
// 下一个时间窗口observer+滑动窗口尺寸
func nextTimeWindow(observer, rollingWindow int64) (next, current int64, canSwitch bool) {
	now := timestamp.Now()
	next = observer + rollingWindow
	if now >= next {
		canSwitch = true
	}
	elapsed := timestamp.SinceZero(now)
	current = elapsed
	return
}

// 获取当前观察点
// 当日0的毫秒数zero + offsetMilliSeconds
func getCurrentObserver(offsetMilliSeconds int64) int64 {
	zero := timestamp.Today()
	return zero + offsetMilliSeconds
}

// RollingOnce 周期性懒加载机制
type RollingOnce struct {
	done            uint32
	m               sync.Mutex
	once            sync.Once     // for ticker
	ticker          *time.Ticker  // 定时器
	currentWindow   int64         // 滑动窗口的毫秒数, 这里默认1天
	currentOffset   atomic.Int64  // 相对于0点整的偏移毫秒数
	currentObserver atomic.Int64  // 当前窗口期起点的毫秒数
	lazyFunc        func()        // 懒加载函数指针
	finished        chan struct{} // 关闭ticker的信号
}

// Close 资源关闭方法
func (o *RollingOnce) Close() {
	// 发送结束信号
	o.finished <- struct{}{}
	close(o.finished)
}

// GetCurrentAnchorPoint 获取当前时间窗口期的锚点
func (o *RollingOnce) GetCurrentAnchorPoint() int64 {
	return o.currentObserver.Load()
}

func (o *RollingOnce) initTicker() {
	// 1. 设置窗口期
	o.currentWindow = rollingWindow
	// 2. 第一步初始化offset, // 偏移默认是常量offsetWindows
	o.currentOffset.CompareAndSwap(0, offsetWindow)
	// 3. 第二步初始化当前时间窗口观察点
	o.currentObserver.CompareAndSwap(0, getCurrentObserver(o.currentOffset.Load()))
	o.finished = make(chan struct{})
	if o.ticker == nil {
		go o.runTicker()
	}
}

// SetOffsetTime 用小时数,分钟数设置滑动窗口的偏移量
func (o *RollingOnce) SetOffsetTime(hour, minute int) {
	offset := timestamp.MillisecondsPerHour * hour
	offset += timestamp.MillisecondsPerMinute * minute
	o.SetOffsetForZero(int64(offset))
}

// SetOffsetForZero 设置时间窗口变化的偏移量
//
//	为非默认9点整重置done预留的功能性方法
func (o *RollingOnce) SetOffsetForZero(offsetMilliSeconds int64) {
	o.currentOffset.Store(offsetMilliSeconds)
	o.updateObserverOfWindow()
}

// 更新窗口期的观察点
func (o *RollingOnce) updateObserverOfWindow() {
	o.currentObserver.Store(getCurrentObserver(o.currentOffset.Load()))
}

// WindowIsExpired 检查当前窗口期的是否过期
func (o *RollingOnce) WindowIsExpired() bool {
	_, _, canSwitch := nextTimeWindow(o.currentObserver.Load(), o.currentWindow)
	return canSwitch
}

func (o *RollingOnce) runTicker() {
	funcName := runtime.FuncName(o.lazyFunc)
	o.ticker = time.NewTicker(100 * time.Millisecond)
	defer o.ticker.Stop()
	for {
		select {
		case <-o.ticker.C:
			// 检查滑动窗口期是否过时
			if o.WindowIsExpired() {
				if runtime.Debug() {
					logger.Infof("RollingOnce[%s]: reset begin", funcName)
				}
				o.Reset()
				if runtime.Debug() {
					logger.Infof("RollingOnce[%s]: reset end", funcName)
				}
			}
		case <-o.finished:
			// 收到结束信号, 退出循环
			return
		}
	}
}

func (o *RollingOnce) Do(f func()) {
	if o.lazyFunc == nil {
		o.lazyFunc = f
	}
	o.once.Do(o.initTicker)
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *RollingOnce) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

// Reset 被动的方式重置初始化done标志
func (o *RollingOnce) Reset() {
	if atomic.LoadUint32(&o.done) == 1 {
		o.resetSlow()
	}
}

func (o *RollingOnce) resetSlow() {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 1 {
		atomic.StoreUint32(&o.done, 0)
		// 重置观察点
		o.updateObserverOfWindow()
	}
}
