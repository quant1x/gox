package coroutine

import (
	"sync"
	"sync/atomic"
)

type PeriodicOnce struct {
	done uint32
	m    sync.Mutex
}

func (o *PeriodicOnce) Do(f func()) {
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
