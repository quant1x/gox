package cache

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	Name string
}

func TestPool(t *testing.T) {
	var pool Pool[TestStruct]
	count := 100
	var t1 TestStruct
	t1.Name = "test"
	pool.Release(&t1)
	for i := 0; i < count; i++ {
		t1 := pool.Acquire()
		fmt.Printf("%d: %p, %+v\n", i, t1, t1)
		t1.Name = fmt.Sprintf("%d", i)
		pool.Release(t1)
	}
}
