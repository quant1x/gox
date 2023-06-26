package main

import (
	"fmt"

	"gitee.com/quant1x/gox/fastqueue"
)

func main() {
	myDisruptor := fastqueue.NewDisruptor(
		fastqueue.WithCapacity(BufferSize),
		fastqueue.WithConsumerGroup(MyConsumer{}))

	go publish(myDisruptor)

	myDisruptor.Read()
}

func publish(myDisruptor fastqueue.Disruptor) {
	for sequence := int64(0); sequence <= Iterations; {
		sequence = myDisruptor.Reserve(Reservations)

		for lower := sequence - Reservations + 1; lower <= sequence; lower++ {
			ringBuffer[lower&BufferMask] = lower
		}

		myDisruptor.Commit(sequence-Reservations+1, sequence)
	}

	_ = myDisruptor.Close()
}

// ////////////////////

type MyConsumer struct{}

func (this MyConsumer) Consume(lower, upper int64) {
	for ; lower <= upper; lower++ {
		message := ringBuffer[lower&BufferMask]
		if message != lower {
			panic(fmt.Errorf("race condition: %d %d", message, lower))
		}
	}
}

const (
	BufferSize   = 1024 * 64
	BufferMask   = BufferSize - 1
	Iterations   = 128 * 1024 * 32
	Reservations = 1
)

var ringBuffer = [BufferSize]int64{}
