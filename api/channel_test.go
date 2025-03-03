package api

import (
	"fmt"
	"testing"
)

func TestChanIsClosed(t *testing.T) {
	// 使用示例
	ch := make(chan int)
	close(ch)
	fmt.Println(ChanIsClosed(ch))   // 输出 true
	fmt.Println(v1ChanIsClosed(ch)) // 输出 true
	fmt.Println(v2ChanIsClosed(ch)) // 输出 true
}
