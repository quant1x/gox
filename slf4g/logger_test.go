package slf4g

import (
	"fmt"
	//"tars/util/logger"
	"testing"
	"time"
)

func XTestLogger(t *testing.T) {
	SetLogPath("/data/logs/test")
	//logger := api.GetLogger("test1")
	//SetConsole()
	for i := 0; i < 2000; i++  {
		Infof("info-%d", i)
		time.Sleep(time.Millisecond * 1)
	}


	Debug("debug")
	Error("error")
	Warn("warn")
	fmt.Println("ok")
	//slf4g.FlushLogger()
	FlushLogger()
}
