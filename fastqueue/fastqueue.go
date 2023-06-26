package fastqueue

import (
	"errors"
	"math/bits"
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
	done       chan bool
}

//// math/bits.OnesCount64
//func bitCount(x int64) int {
//	// HD, Figure 5-2
//	x = x - ((x >> 1) & 0x55555555)
//	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
//	x = (x + (x >> 4)) & 0x0f0f0f0f
//	x = x + (x >> 8)
//	x = x + (x >> 16)
//	x = x & 0x3f
//	return int(x)
//}

func NewQueue[T any](max ...int) *Queue[T] {
	buffSize := int64(kBufferSize)
	if len(max) > 0 {
		buffSize = int64(max[0])
	}
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

func (q *Queue[T]) Wait() {
	q.disruptor.Read()
}

func (q *Queue[T]) SetEvent(event func(data []T)) {
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

func (q *Queue[T]) pushOne(data T) {
	kReservations := int64(1)
	sequence := q.disruptor.Reserve(kReservations)
	for lower := sequence - kReservations + 1; lower <= sequence; lower++ {
		q.ringBuffer[lower&q.indexMask] = data
	}
	q.disruptor.Commit(sequence-kReservations+1, sequence)
}

func (q *Queue[T]) Push(data ...T) {
	num := len(data)
	if num == 0 {
		return
	}
	for _, v := range data {
		q.pushOne(v)
	}
}
