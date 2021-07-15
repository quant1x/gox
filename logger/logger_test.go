package logger

import (
	"fmt"
	"github.com/mymmsc/gox/gls"
	"github.com/mymmsc/gox/mdc"
	"github.com/mymmsc/gox/util/uuid"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	SetLogPath("/opt/logs/test")
	u1 := uuid.NewV4()
	defer gls.DeleteGls(gls.GoID())
	mdc.Set(mdc.APP_TRACEID, u1.String())
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
	mdc.Remove(mdc.APP_TRACEID)
}
