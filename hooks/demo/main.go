package main

import (
	"context"
	"fmt"
	"time"

	"gitee.com/quant1x/gox/hooks"
)

func main() {
	//logger.SetLevel(logger.INFO)
	hm := hooks.NewHookManager()

	// 注册高优先级 Hook（如关闭数据库）
	hm.Register(func(ctx context.Context) error {
		fmt.Println("Closing database...")
		time.Sleep(1 * time.Second)
		return nil
	}, hooks.WithPriority(hooks.HighPriority))

	// 注册低优先级 Hook（如清理临时文件）
	hm.Register(func(ctx context.Context) error {
		fmt.Println("Cleaning temp files...")
		time.Sleep(500 * time.Millisecond)
		return nil
	}, hooks.WithPriority(hooks.LowPriority))

	// 启动监听
	hm.Listen()

	// 主程序逻辑
	fmt.Println("Server running, press Ctrl+C to exit")
	hm.Wait() // 阻塞直到退出信号
}
