package api

import (
	"time"
)

type FileStat struct {
	CreationTime   time.Time
	LastAccessTime time.Time
	LastWriteTime  time.Time
}

// 把秒级的时间戳转为time格式
func secondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

func NanosecondToTime(nanoseconds int64) time.Time {
	seconds := int64(time.Second)
	sec := nanoseconds / seconds
	nsec := nanoseconds % seconds
	return time.Unix(sec, nsec)
}
