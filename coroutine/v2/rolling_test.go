package coroutine

import (
	"fmt"
	"testing"
	"time"
)

// 测试用例示例
func Test2HourWindow(t *testing.T) {
	task := func() { fmt.Println("EXECUTE") }
	ro, err := NewRollingOnce(
		2*60*60*1000, // 2小时
		30*60*1000,   // 30分钟偏移
		task,
	)
	if err != nil {
		t.Fatal(err)
	}

	// 模拟时间流逝验证窗口计算
	ro.observer.Store(time.Date(2023, 1, 1, 9, 30, 0, 0, time.Local).UnixMilli())
	ro.updateObserver()
	next := time.UnixMilli(ro.observer.Load())
	if next.Hour() != 11 || next.Minute() != 30 { // 期望下次执行在11:30
		t.Fatalf("窗口计算错误，实际时间: %v", next)
	}
}
