package coroutine

import (
	"fmt"
	"gitee.com/quant1x/gox/concurrent"
	"strconv"
	"testing"
	"time"
)

var (
	rollineOnce RollingOnce
	cache1      = concurrent.NewHashMap[string, int]()
)

func lazyCacheInit() {
	for i := 0; i < 5; i++ {
		k := strconv.Itoa(i)
		//cache[k] = i
		cache1.Set(k, i)
		fmt.Println("reset", k, "=>", i)
	}
}

func cacheGetInt(key string) (int, bool) {
	rollineOnce.Do(lazyCacheInit)
	v, ok := cache1.Get(key)
	return v, ok
}

func cacheSetInt(key string, value int) {
	rollineOnce.Do(lazyCacheInit)
	cache1.Set(key, value)
}

func TestRollingOnce(t *testing.T) {
	rwCount := 1000
	producer := func() {
		for i := 0; i < rwCount; i++ {
			k := strconv.Itoa(i % 5)
			cacheSetInt(k, i)
			fmt.Println(k, "=>", i)
			time.Sleep(time.Millisecond * 10)
		}
	}
	reader := func() {
		for i := 0; i < rwCount; i++ {
			k := strconv.Itoa(i % 5)
			v, ok := cacheGetInt(k)
			fmt.Println(v, "<=", i, ":", ok)
			_ = v
			_ = ok
			time.Sleep(time.Millisecond * 10)
		}
	}

	go producer()
	go reader()
	count := 60
	for i := 0; i < count; i++ {
		//once.Reset()
		fmt.Println("--------------------")
		time.Sleep(time.Second)
	}
}

func Test_defaultTimeWindow(t *testing.T) {
	observer := currentObserver()
	fmt.Println(observer)
	a, b, c := defaultTimeWindow(observer, rollingWindow)
	fmt.Println(a, b, c)
}
