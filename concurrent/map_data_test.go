package concurrent

import (
	"math/rand"
	"strconv"
	"sync"
)

var (
	testTreeMapOnce sync.Once
	testTreemap     = NewTreeMap[string, int]()
)

const (
	testCount = 10000
)

// 初始化基准测试数据
func init() {
	for i := 0; i < testCount; i++ {
		key := strconv.Itoa(i)
		value := rand.Int()
		testTreemap.Put(key, value)
	}
}
