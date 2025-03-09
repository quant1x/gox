package coroutine

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrChannelClosed = errors.New("channel closed")
	ErrBufferFull    = errors.New("buffer full")
)

type Channel[E any] struct {
	data      chan E
	closeOnce sync.Once
	closed    atomic.Bool
	num       atomic.Int64
}

func NewChannel[E any](bufSize int) *Channel[E] {
	return &Channel[E]{
		data: make(chan E, bufSize),
	}
}

func (c *Channel[E]) Push(v E) error {
	if c.closed.Load() {
		return ErrChannelClosed
	}

	select {
	case c.data <- v:
		c.num.Add(1)
		return nil
	default:
		return ErrBufferFull
	}
}

// SafePush 带超时的安全推送
func (c *Channel[E]) SafePush(v E, timeout time.Duration) error {
	if c.closed.Load() {
		return ErrChannelClosed
	}

	select {
	case c.data <- v:
		c.num.Add(1)
		return nil
	case <-time.After(timeout):
		return ErrBufferFull
	}
}

func (c *Channel[E]) Pop() (E, bool) {
	v, ok := <-c.data
	if ok {
		c.num.Add(-1)
	}
	return v, ok
}

func (c *Channel[E]) Close() {
	c.closeOnce.Do(func() {
		close(c.data)
		c.closed.Store(true)
	})
}

func (c *Channel[E]) Len() int {
	return int(c.num.Load())
}

func (c *Channel[E]) Cap() int {
	return cap(c.data)
}
