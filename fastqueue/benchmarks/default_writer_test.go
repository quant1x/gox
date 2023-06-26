package benchmarks

import (
	"log"
	"math"
	"testing"

	"gitee.com/quant1x/gox/fastqueue"
)

func BenchmarkWriterReserve(b *testing.B) {
	read, written := fastqueue.NewCursor(), fastqueue.NewCursor()
	writer := fastqueue.NewWriter(written, read, 1024)
	iterations := int64(b.N)
	b.ReportAllocs()
	b.ResetTimer()

	for i := int64(0); i < iterations; i++ {
		sequence := writer.Reserve(1)
		read.Store(sequence)
	}
}
func BenchmarkWriterNextWrapPoint(b *testing.B) {
	read, written := fastqueue.NewCursor(), fastqueue.NewCursor()
	writer := fastqueue.NewWriter(written, read, 1024)
	iterations := int64(b.N)
	b.ReportAllocs()
	b.ResetTimer()

	read.Store(math.MaxInt64)
	for i := int64(0); i < iterations; i++ {
		writer.Reserve(1)
	}
}
func BenchmarkWriterCommit(b *testing.B) {
	writer := fastqueue.NewWriter(fastqueue.NewCursor(), nil, 1024)
	iterations := int64(b.N)
	b.ReportAllocs()
	b.ResetTimer()

	for i := int64(0); i < iterations; i++ {
		writer.Commit(i, i)
	}
}

func BenchmarkWriterReserveOneSingleConsumer(b *testing.B) {
	benchmarkSequencerReservations(b, ReserveOne, SampleConsumer{})
}
func BenchmarkWriterReserveManySingleConsumer(b *testing.B) {
	benchmarkSequencerReservations(b, ReserveMany, SampleConsumer{})
}
func BenchmarkWriterReserveOneMultipleConsumers(b *testing.B) {
	benchmarkSequencerReservations(b, ReserveOne, SampleConsumer{}, SampleConsumer{})
}
func BenchmarkWriterReserveManyMultipleConsumers(b *testing.B) {
	benchmarkSequencerReservations(b, ReserveMany, SampleConsumer{}, SampleConsumer{})
}
func benchmarkSequencerReservations(b *testing.B, count int64, consumers ...fastqueue.Consumer) {
	iterations := int64(b.N)

	myDisruptor := build(consumers...)

	go func() {
		b.ReportAllocs()
		b.ResetTimer()

		var sequence int64 = -1
		for sequence < iterations {
			sequence = myDisruptor.Reserve(count)
			for i := sequence - (count - 1); i <= sequence; i++ {
				ringBuffer[i&RingBufferMask] = i
			}
			myDisruptor.Commit(sequence, sequence)
		}

		_ = myDisruptor.Close()
	}()

	myDisruptor.Read()
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SampleConsumer struct{}

func (this SampleConsumer) Consume(lower, upper int64) {
	var message int64
	for lower <= upper {
		message = ringBuffer[lower&RingBufferMask]
		if message != lower {
			log.Panicf("race condition: Sequence: %d, Message: %d", lower, message)
		}
		lower++
	}
}

func build(consumers ...fastqueue.Consumer) fastqueue.Disruptor {
	return fastqueue.NewDisruptor(
		fastqueue.WithCapacity(RingBufferSize),
		fastqueue.WithConsumerGroup(consumers...))
}

const (
	RingBufferSize = 1024 * 64
	RingBufferMask = RingBufferSize - 1
	ReserveOne     = 1
	ReserveMany    = 16
)

var ringBuffer = [RingBufferSize]int64{}
