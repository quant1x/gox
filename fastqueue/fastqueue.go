package fastqueue

import (
	"errors"
	"math/bits"
	"sync"
)

const (
	BUFFER_PAD  = 32 // java版本有这个pad, 作用不详
	kBufferSize = 1024 * 64
	// kBufferMask   = kBufferSize - 1
	// kReservations = 1
)

var (
	errBufferSizeNotPower2 = errors.New("bufferSize must be a power of 2")
)

type Queue[T any] struct {
	disruptor  Disruptor
	ringBuffer []T
	bufferSize int64
	indexMask  int64
	onData     func([]T)
	waitGroup  sync.WaitGroup
	//done       chan bool
}

func highestOneBit(x uint64) uint64 {
	origin := x
	isPower2 := ((origin & (origin - 1)) == 0) && (origin != 0)
	if isPower2 {
		return origin
	}
	// HD, Figure 3-1
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	//x |= (x >> 64)
	//x = x - (x >> 1)
	x = x & ^(x >> 1)
	if x < origin {
		x <<= 1
	}
	if x == 0 {
		x = 1 << 63
	}
	return x
}

func NewQueue[T any](max ...int) *Queue[T] {
	buffSize := int64(kBufferSize)
	if len(max) > 0 {
		buffSize = int64(max[0])
	}
	buffSize = int64(highestOneBit(uint64(buffSize)))
	if bits.OnesCount64(uint64(buffSize)) != 1 {
		panic(errBufferSizeNotPower2)
	}
	queue := &Queue[T]{}
	queue.bufferSize = buffSize
	queue.indexMask = queue.bufferSize - 1
	//queue.ringBuffer = make([]T, queue.bufferSize+2*BUFFER_PAD)
	queue.ringBuffer = make([]T, queue.bufferSize)
	queue.disruptor = NewDisruptor(
		WithCapacity(buffSize),
		WithConsumerGroup(queue),
	)
	return queue
}

func (q *Queue[T]) Finish() {
	_ = q.disruptor.Close()
}

func (q *Queue[T]) Wait() {
	go func() {
		q.waitGroup.Wait()
		q.Finish()
	}()
	q.disruptor.Read()
}

func (q *Queue[T]) SetReadEvent(event func(data []T)) {
	q.onData = event
}

func (q *Queue[T]) Consume(lower, upper int64) {
	list := []T{}
	for ; lower <= upper; lower++ {
		message := q.ringBuffer[lower&q.indexMask]
		list = append(list, message)
	}
	if len(list) > 0 {
		q.onData(list)
	}
}

func queuePushMessage[T any](queue *Queue[T], data T) {
	defer queue.waitGroup.Done()
	kReservations := int64(1)
	sequence := queue.disruptor.Reserve(kReservations)
	for lower := sequence - kReservations + 1; lower <= sequence; lower++ {
		queue.ringBuffer[lower&queue.indexMask] = data
	}
	queue.disruptor.Commit(sequence-kReservations+1, sequence)
}

func (q *Queue[T]) pushOne_old(data T) {
	kReservations := int64(1)
	sequence := q.disruptor.Reserve(kReservations)
	for lower := sequence - kReservations + 1; lower <= sequence; lower++ {
		q.ringBuffer[lower&q.indexMask] = data
	}
	q.disruptor.Commit(sequence-kReservations+1, sequence)
}

func (q *Queue[T]) PushOne(data T) {
	q.waitGroup.Add(1)
	go queuePushMessage(q, data)
}

func (q *Queue[T]) Push(data ...T) {
	num := len(data)
	if num == 0 {
		return
	}
	for _, v := range data {
		q.PushOne(v)
	}
}
