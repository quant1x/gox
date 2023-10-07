package concurrent

import (
	"fmt"
	"testing"
)

func TestTreeMap(t *testing.T) {
	treemap := NewTreeMap[string, int]()
	treemap.Put("a", 10)
	treemap.Put("1a", 12)
	v, ok := treemap.Get("a")
	fmt.Println(v, ok)
	v, ok = treemap.Get("a1")
	fmt.Println(v, ok)

	treemap.Each(func(key string, value int) {
		fmt.Println(key, value)
	})
}
