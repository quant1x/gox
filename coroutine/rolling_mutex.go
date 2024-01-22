package coroutine

import (
	"sync"
	"sync/atomic"
	"time"
)

const (
	onceInitTime    = "09:00:00"
	onceDefaultDate = "1970-01-01"
)

// RollingMutex 按指定rolling策略加锁, 指定周期内只加载一次
//
//	滑动窗口锁, 窗口期内只初始化一次, 目前只支持1天切换
//
// Deprecated: 不推荐, 建议使用 RollingOnce
type RollingMutex struct {
	m    sync.Mutex
	date string
	done uint32
}

// 校对当前日期
func (o *RollingMutex) proofreadCurrentDate() (currentDate string) {
	currentDate = o.date
	if currentDate < onceDefaultDate {
		currentDate = onceDefaultDate
	}
	now := time.Now()
	timestamp := now.Format(time.TimeOnly)
	if timestamp >= onceInitTime {
		currentDate = now.Format(time.DateOnly)
	}
	return currentDate
}

func (o *RollingMutex) Do(f func(), today ...func() (newDate string)) {
	getToday := o.proofreadCurrentDate
	if len(today) > 0 {
		getToday = today[0]
	}
	if getToday != nil {
		currentDate := getToday()
		if atomic.LoadUint32(&o.done) == 1 && currentDate > o.date {
			o.doReset(currentDate)
		}
	}
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *RollingMutex) doReset(currentDate string) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 1 && currentDate > o.date {
		atomic.StoreUint32(&o.done, 0)
		o.date = currentDate
	}
}

func (o *RollingMutex) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func (o *RollingMutex) Reset() {
	atomic.StoreUint32(&o.done, 0)
}

func (o *RollingMutex) Date() string {
	return o.date
}
