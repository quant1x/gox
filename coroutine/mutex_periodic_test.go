package coroutine

import (
	"fmt"
	"gitee.com/quant1x/gox/concurrent"
	"strconv"
	"testing"
	"time"
)

var (
	once PeriodicOnce
	//cache = map[string]int{}
	//cache = cmap.NewStringMap[int]()
	cache = concurrent.NewHashMap[string, int]()
)

func lazyInit() {
	for i := 0; i < 5; i++ {
		k := strconv.Itoa(i)
		//cache[k] = i
		cache.Set(k, i)
	}
}

func getInt(key string) (int, bool) {
	once.Do(lazyInit)
	//v, ok := cache[key]
	v, ok := cache.Get(key)
	return v, ok
}
func setInt(key string, value int) {
	once.Do(lazyInit)
	//cache[key] = value
	cache.Set(key, value)
}

func TestPeriodicOnce(t *testing.T) {
	rwCount := 1000
	producer := func() {
		for i := 0; i < rwCount; i++ {
			k := strconv.Itoa(i % 5)
			setInt(k, i)
			time.Sleep(time.Millisecond * 10)
		}
	}
	reader := func() {
		for i := 0; i < rwCount; i++ {
			k := strconv.Itoa(i % 5)
			v, ok := getInt(k)
			fmt.Println(v, ok)
			_ = v
			_ = ok
			time.Sleep(time.Millisecond * 10)
		}
	}

	go producer()
	go reader()
	count := 30
	for i := 0; i < count; i++ {
		once.Reset()
		fmt.Println("--------------------")
		time.Sleep(time.Second)
	}
}
