package coroutine

import (
	"gitee.com/quant1x/gox/cron"
	"gitee.com/quant1x/gox/logger"
	"sync"
	"sync/atomic"
	"time"
)

const (
	cronPreMinute    = "*/1 * * * *"
	cronPreSecond    = "0/1 * * * * ?"
	periodicInitTime = "09:00:00"
)

// PeriodicOnce 周期性懒加载机制
type PeriodicOnce struct {
	done  uint32
	m     sync.Mutex
	once  sync.Once
	timer *cron.Cron
	date  string
}

func (o *PeriodicOnce) Do(f func()) {
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
		o.date = onceDefaultDate
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
		o.timer = cron.New(cron.WithSeconds())
		id, err := o.timer.AddFunc(cronPreSecond, func() {
			if o.isExpired() {
				o.Reset()
			}
		})
		logger.Info(id, err)
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
	}
}
