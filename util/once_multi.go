package util

import (
	"sync"
	"sync/atomic"
)

type MultiOnce struct {
	done uint32
	m    sync.Mutex
	date string
}

func (o *MultiOnce) Do(f func(), today ...func() (newDate string)) {
	var getToday func() (newDate string)
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
