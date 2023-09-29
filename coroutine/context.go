package coroutine

import (
	"context"
	"sync"
)

var (
	globalOnce    sync.Once
	globalContext context.Context    = nil
	globalCancel  context.CancelFunc = nil
)

func initContext() {
	globalContext, globalCancel = context.WithCancel(context.Background())
}

// Context 获取全局顶层context
func Context() context.Context {
	globalOnce.Do(initContext)
	return globalContext
}

// Shutdown 关闭应用程序, 通知所有协程退出
func Shutdown() {
	globalOnce.Do(initContext)
	if globalCancel != nil {
		globalCancel()
	}
}

func GetContextWithCancel() (context.Context, context.CancelFunc) {
	globalOnce.Do(initContext)
	ctx, cancel := context.WithCancel(globalContext)
	return ctx, cancel
}
