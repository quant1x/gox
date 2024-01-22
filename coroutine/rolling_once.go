package coroutine

import (
	"gitee.com/quant1x/gox/logger"
	"gitee.com/quant1x/gox/runtime"
	"gitee.com/quant1x/gox/timestamp"
	"sync"
	"sync/atomic"
	"time"
)

const (
	rollingAnchorPoint = 0 // 滑动窗口的锚点毫秒数
	rollingWindow      = timestamp.MillisecondsPerDay
	offsetWindow       = timestamp.MillisecondsPerHour * 9 // 每天9点整
)

// 默认的时间窗口
func defaultTimeWindow(observer, rollingWindow int64) (next, current int64, canSwitch bool) {
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
func currentObserver(offsetMilliSeconds int64) int64 {
	zero := timestamp.Today()
	return zero + offsetMilliSeconds
}

// RollingOnce 周期性懒加载机制
type RollingOnce struct {
	done     uint32
	m        sync.Mutex
	once     sync.Once     // for ticker
	ticker   *time.Ticker  // 定时器
	window   int64         // TODO: 暂时未起作用, 滑动窗口的毫秒数, 这里是1天
	offset   int64         // 距离0点整的偏移毫秒数
	observer int64         // 当前窗口期的毫秒数
	lazyFunc func()        // 懒加载函数指针
	finished chan struct{} // 关闭ticker的信号
}

func (o *RollingOnce) Close() {
	// 发送结束信号
	o.finished <- struct{}{}
	// TODO: 不确定实时性, 暂时屏蔽close操作
	//close(o.finished)
}

func (o *RollingOnce) initTicker() {
	// 1. 第一步初始化offset
	if o.offset == 0 {
		// 偏移默认是常量offsetWindows
		o.offset = offsetWindow
	}
	// 2. 第二步初始化当前时间窗口
	if o.observer == 0 {
		// 首次设置当前时间窗口的观察时间戳
		o.observer = currentObserver(o.offset)
	}
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
	o.m.Lock()
	defer o.m.Unlock()
	o.offset = offsetMilliSeconds
	o.observer = currentObserver(o.offset)
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

func (o *RollingOnce) windowIsExpired() bool {
	_, _, canSwitch := defaultTimeWindow(o.observer, o.offset)
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
			if o.windowIsExpired() {
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
			break
		}
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
		o.observer = currentObserver(o.offset)
	}
}
