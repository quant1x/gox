package logger

import (
	"fmt"
	"os"
)

// 等待进程结束信号
func waitForStop(ch chan os.Signal) {
	s := <-ch
	for _, bw := range bws {
		err := bw.Stop()
		if err != nil {
			Errorf("zapcore.BufferedWriteSyncer stop error: %v", err)
		}
	}
	Infof("exit sign, [%+v]", s)
	fmt.Println("exit", s)
	if logger != nil {
		_ = logger.Desugar().Sync()
	}
	os.Exit(0)
}
