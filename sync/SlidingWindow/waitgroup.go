package SlidingWindow

import (
	"context"
	"math"
	"sync"
)

// SlidingWindow has the same role and close to the
// same API as the Golang sync.WaitGroup but adds a limit of
// the amount of goroutines started concurrently.
type SlidingWindow struct {
	Size int

	current chan struct{}
	wg      sync.WaitGroup
}

// New creates a SlidingWindow.
// The limit parameter is the maximum amount of
// goroutines which can be started concurrently.
func New(limit int) SlidingWindow {
	size := math.MaxInt32 // 2^32 - 1
	if limit > 0 {
		size = limit
	}
	return SlidingWindow{
		Size: size,

		current: make(chan struct{}, size),
		wg:      sync.WaitGroup{},
	}
}

// Add increments the internal WaitGroup counter.
// It can be blocking if the limit of spawned goroutines
// has been reached. It will stop blocking when Done is
// been called.
//
// See sync.WaitGroup documentation for more information.
func (s *SlidingWindow) Add() {
	s.AddWithContext(context.Background())
}

// AddWithContext increments the internal WaitGroup counter.
// It can be blocking if the limit of spawned goroutines
// has been reached. It will stop blocking when Done is
// been called, or when the context is canceled. Returns nil on
// success or an error if the context is canceled before the lock
// is acquired.
//
// See sync.WaitGroup documentation for more information.
func (s *SlidingWindow) AddWithContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.current <- struct{}{}:
		break
	}
	s.wg.Add(1)
	return nil
}

// Done decrements the SlidingWindow counter.
// See sync.WaitGroup documentation for more information.
func (s *SlidingWindow) Done() {
	<-s.current
	s.wg.Done()
}

// Wait blocks until the SlidingWindow counter is zero.
// See sync.WaitGroup documentation for more information.
func (s *SlidingWindow) Wait() {
	s.wg.Wait()
}
