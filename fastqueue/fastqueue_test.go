package fastqueue

import (
	"fmt"
	"math/bits"
	"testing"
)

const (
	//tIterations = 128 * 1024 * 32
	tIterations = 128
)

func TestQueue_Consume(t *testing.T) {
	queue := NewQueue[string](tIterations - 1)
	queue.SetReadEvent(func(data []string) {
		for _, v := range data {
			//time.Sleep(time.Second * 1)
			fmt.Println(v)
		}
	})
	for i := 0; i < tIterations*2; i++ {
		message := fmt.Sprintf("%d", i)
		go queue.Push(message)
	}

	//time.Sleep(time.Second * 10)
	//go queue.PushOne("128")
	//queue.Finish()
	queue.Wait()
	//time.Sleep(time.Second * 1)
}

// #define APR_ALIGN(size, boundary)    (((size) + ((boundary) - 1)) & ~((boundary) - 1))
//
//	将size往上"取整到"大于size的最小的boundary的倍数
func AlignUint(x uint64) uint64 {
	const boundary uint64 = 2
	y := ((x) + ((boundary) - 1)) & ^((boundary) - 1)
	return y
}

func Test_bitCount(t *testing.T) {
	fmt.Println(highestOneBit(1))
	fmt.Println(highestOneBit(2))
	fmt.Println(highestOneBit(3))
	fmt.Println(highestOneBit(4))
	fmt.Println(highestOneBit(126))
	fmt.Println(highestOneBit(127))
	fmt.Println(highestOneBit(128))
	fmt.Println(highestOneBit(0xffffffffffffffff))
	var x uint64 = 3
	xl := bits.Len64(x)
	if xl%2 != 0 {
		xl++
	}
	x = 1 << (xl)
	fmt.Println(bits.OnesCount64(x))
	fmt.Println(bits.OnesCount64(128))
	fmt.Println(bits.OnesCount64(256))

	fmt.Println(x, "len =", bits.Len64(x))
	x = 128
	fmt.Println(x, "len =", bits.Len64(x))
}
