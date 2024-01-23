package cache

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	key := "1"
	var t1 TestStruct
	t1.Name = "test"
	var m Map[string, TestStruct]
	m.Put(key, t1)
	t2, ok := m.Get(key)
	fmt.Println(t2, ok)
	key = "2"
	t3, ok := m.Get(key)
	fmt.Println(t3, ok)
}
