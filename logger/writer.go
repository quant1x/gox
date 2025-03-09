package logger

import (
	"fmt"
	"gitee.com/quant1x/gox/api"
	"os"
)

const (
	DAY DateType = iota
	HOUR
)

type LogWriter interface {
	Write(v []byte)
	NeedPrefix() bool
}

// 重新打开文件
func reopenFile(name string, currFile **os.File, openTime *int64) {
	*openTime = timeRotate.GetUnixTime()
	if *currFile != nil {
		_ = (*currFile).Close()
	}
	of, err := os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, api.CACHE_FILE_MODE)
	if err == nil {
		*currFile = of
	} else {
		fmt.Println("open log file error", err)
	}
}
