package runtime

import (
	"fmt"
	"gitee.com/quant1x/gox/logger"
	"runtime/debug"
)

// CatchPanic 捕获panic
func CatchPanic() {
	if err := recover(); err != nil {
		s := string(debug.Stack())
		fmt.Printf("\nerr=%v, stack=%s\n", err, s)
		logger.Fatalf("%s 异常: %+v", ApplicationName(), err)
	}
}

// IgnorePanic 通用捕获panic, 忽略异常, 继续执行
func IgnorePanic() {
	if err := recover(); err != nil {
		s := string(debug.Stack())
		logger.Errorf("err=%v, stack=%s", err, s)
	}
}
