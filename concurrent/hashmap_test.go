package concurrent

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {
	m := v1NewHashmap[string, string]()
	m.Put("a", "1")
	v, ok := m.Get("a")
	fmt.Println(v, ok)
}
