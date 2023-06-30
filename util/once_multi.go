package util

import (
	"sync"
	"sync/atomic"
	"time"
)

const (
	onceInitTime    = "09:00:00"
	onceDefaultDate = "1970-01-01"
)

type MultiOnce struct {
	done uint32
	m    sync.Mutex
	date string
}

// 校对当前日期
func proofreadCurrentDate() (currentDate string) {
	now := time.Now()
	timestamp := now.Format(time.TimeOnly)
	if timestamp >= onceInitTime {
		currentDate = now.Format(time.DateOnly)
	}
	return onceDefaultDate
}

func (o *MultiOnce) Do(f func(), today ...func() (newDate string)) {
	getToday := proofreadCurrentDate
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

func (o *MultiOnce) doReset(currentDate string) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 1 && currentDate > o.date {
		o.m.Lock()
		defer o.m.Unlock()
		atomic.StoreUint32(&o.done, 0)
		o.date = currentDate
	}
}

func (o *MultiOnce) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func (o *MultiOnce) Reset() {
	atomic.StoreUint32(&o.done, 0)
}

func (o *MultiOnce) Date() string {
	return o.date
}
