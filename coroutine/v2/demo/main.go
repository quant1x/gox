package main

import (
	"fmt"
	"gitee.com/quant1x/gox/coroutine/v2"
	"gitee.com/quant1x/gox/timestamp"
	"time"
)

func main() {
	// 创建每天9:30执行的任务
	ro, err := coroutine.NewRollingOnce(
		timestamp.MillisecondsPerDay,
		9*timestamp.MillisecondsPerHour+30*timestamp.MillisecondsPerMinute,
		func() {
			fmt.Println("执行每日数据清理...")
		},
	)
	if err != nil {
		panic(err)
	}
	defer ro.Close()

	// 模拟多次调用
	for i := 0; i < 10; i++ {
		ro.Do()
		time.Sleep(1 * time.Second)
	}
}
