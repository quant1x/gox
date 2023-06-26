package fastqueue

import (
	"fmt"
	"math/bits"
	"testing"
	"time"
)

const (
	//tIterations = 128 * 1024 * 32
	tIterations = 128
)

func TestQueue_Consume(t *testing.T) {
	queue := NewQueue[string](tIterations)
	queue.SetEvent(func(data []string) {
		for _, v := range data {
			fmt.Println(v)
		}
	})
	//wg := sync.WaitGroup{}
	for i := 0; i < tIterations; i++ {
		//wg.Add(1)
		message := fmt.Sprintf("%d", i)
		queue.pushOne(message)
	}
	//wg.Wait()
	queue.Wait()
	time.Sleep(time.Second * 10000)
}

func Test_bitCount(t *testing.T) {
	fmt.Println(bits.OnesCount64(127))
	fmt.Println(bits.OnesCount64(128))
	fmt.Println(bits.OnesCount64(256))
}
