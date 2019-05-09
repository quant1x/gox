package logger

import (
	"fmt"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	SetLogPath("/data/logs/test")
	//logger := api.GetLogger("test1")
	//SetConsole()
	for i := 0; i < 200; i++ {
		Infof("info-%d", i)
		time.Sleep(time.Millisecond * 1)
	}
	Infof("测试中文\n")
	Debug("debug")
	Error("error")
	Warn("warn")
	Info("测试中文")
	fmt.Println("ok")
	//logger.FlushLogger()
	FlushLogger()
}
