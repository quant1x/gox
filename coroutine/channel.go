package coroutine

import (
	"sync"
	"sync/atomic"
)

// Channel 多协程安全的channel
type Channel[E any] struct {
	data chan E
	m    sync.Mutex
	num  atomic.Int64
	done atomic.Uint32
}

func (c *Channel[E]) Push(v E) {
	if c.done.Load() == 1 {
		return
	}
	c.data <- v
	c.num.Add(1)
}

func (c *Channel[E]) Pop() E {
	v := <-c.data
	c.num.Add(-1)
	return v
}

func (c *Channel[E]) Close() {
	if c.done.Load() != 0 {
		return
	}
	c.m.Lock()
	defer c.m.Unlock()
	if c.done.Load() == 0 {
		c.done.Store(1)
	}
}
