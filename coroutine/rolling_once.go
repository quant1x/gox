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
	rollingAnchorPoint = 0
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
func currentObserver() int64 {
	zero := timestamp.Today()
	return zero + offsetWindow
}

// RollingOnce 周期性懒加载机制
type RollingOnce struct {
	done     uint32
	m        sync.Mutex
	once     sync.Once // for ticker
	ticker   *time.Ticker
	window   int64 // 滑动窗口的毫秒数
	offset   int64 // 距离0点整的偏移毫秒数
	observer int64
	lazyFunc func()
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

func (o *RollingOnce) isExpired() bool {
	_, _, canSwitch := defaultTimeWindow(o.observer, rollingWindow)
	return canSwitch
}

func (o *RollingOnce) runTicker() {
	funcName := runtime.FuncName(o.lazyFunc)
	o.ticker = time.NewTicker(100 * time.Millisecond)
	defer o.ticker.Stop()
	for {
		select {
		case <-o.ticker.C:
			if o.isExpired() {
				if runtime.Debug() {
					logger.Infof("RollingOnce[%s]: reset begin", funcName)
				}
				o.Reset()
				if runtime.Debug() {
					logger.Infof("RollingOnce[%s]: reset end", funcName)
				}
			}
		}
	}
}

func (o *RollingOnce) initTicker() {
	if o.observer == 0 {
		o.observer = currentObserver()
	}
	if o.ticker == nil {
		go o.runTicker()
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
		o.observer = currentObserver()
	}
}
