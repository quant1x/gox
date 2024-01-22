package coroutine

import (
	"gitee.com/quant1x/gox/cron"
	"gitee.com/quant1x/gox/logger"
	"gitee.com/quant1x/gox/runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	cronPreMinute       = "*/1 * * * *"
	cronPreSecond       = "*/1 * * * * ?"
	cronDefaultInterval = "@every 1s"
	periodicInitTime    = "09:00:00"
)

// PeriodicOnce 周期性懒加载机制
//
// Deprecated: 推荐 RollingOnce [wangfeng on 2024/1/22 10:33]
type PeriodicOnce struct {
	done     uint32
	m        sync.Mutex
	once     sync.Once
	timer    *cron.Cron
	date     string
	lazyFunc func()
}

func (o *PeriodicOnce) Do(f func()) {
	if o.lazyFunc == nil {
		o.lazyFunc = f
	}
	o.once.Do(o.initTimer)
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *PeriodicOnce) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func (o *PeriodicOnce) isExpired() bool {
	currentDate := o.date
	if currentDate < onceDefaultDate {
		currentDate = onceDefaultDate
	}
	now := time.Now()
	timestamp := now.Format(time.TimeOnly)
	if timestamp >= periodicInitTime {
		currentDate = now.Format(time.DateOnly)
	}
	if currentDate > o.date {
		return true
	}
	return false
}

func (o *PeriodicOnce) initTimer() {
	if o.timer == nil {
		funcName := runtime.FuncName(o.lazyFunc)
		o.timer = cron.New(cron.WithSeconds())
		_, err := o.timer.AddFuncWithSkipIfStillRunning(cronDefaultInterval, func() {
			if o.isExpired() {
				if runtime.Debug() {
					logger.Infof("PeriodicOnce[%s]: reset begin", funcName)
				}
				o.Reset()
				if runtime.Debug() {
					logger.Infof("PeriodicOnce[%s]: reset end", funcName)
				}
			}
		})
		if err == nil {
			o.timer.Start()
		}
	}
}

// Reset 被动的方式重置初始化done标志
func (o *PeriodicOnce) Reset() {
	if atomic.LoadUint32(&o.done) == 1 {
		o.resetSlow()
	}
}

func (o *PeriodicOnce) resetSlow() {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 1 {
		atomic.StoreUint32(&o.done, 0)
		// 重置日期
		now := time.Now()
		o.date = now.Format(time.DateOnly)
	}
}
