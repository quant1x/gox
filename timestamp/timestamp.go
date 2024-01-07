// Package timestamp 本地时间戳相关功能
package timestamp

import (
	_ "runtime"
	"time"
	_ "unsafe" // for go:linkname
)

const (
	secondsPerMinute      = 60
	secondsPerHour        = 60 * secondsPerMinute
	secondsPerDay         = 24 * secondsPerHour
	millisecondsPerSecond = 1000
	millisecondsPerMinute = secondsPerMinute * millisecondsPerSecond
	millisecondsPerHour   = secondsPerHour * millisecondsPerSecond
	millisecondsPerDay    = secondsPerDay * millisecondsPerSecond
)

//go:linkname now time.now
func now() (sec int64, nsec int32, mono int64)

// 调用公开结构的私有方法
//
//go:linkname abstime time.Time.abs
func abstime(t time.Time) uint64

var (
	// 获取偏移的秒数
	zoneName, offsetInSecondsEastOfUTC = time.Now().Zone()
	_                                  = zoneName
	// UTC到本地的偏移秒数
	utcToLocal = int64(offsetInSecondsEastOfUTC)
	// 本地到UTC的偏移秒数
	localToUTC = -utcToLocal
	// 偏移的毫秒数
	//offsetMilliseconds = offsetInSecondsEastOfUTC * millisecondsPerSecond
)

// Now 获取本地当前的时间戳, 毫秒数
//
//	UTC 转 local
func Now() int64 {
	sec, nsec, _ := now()
	sec += int64(offsetInSecondsEastOfUTC)
	milli := sec*millisecondsPerSecond + int64(nsec)/1e6%millisecondsPerSecond
	return milli
}

// Today 获取当日0点整的时间戳, 毫秒数
//
//	UTC 转 local
func Today() int64 {
	now := Now()
	millisecondsOfToday := now - now%millisecondsPerDay
	return millisecondsOfToday
}

// Timestamp 获取time.Time的本地毫秒数
//
//	UTC 转 local
func Timestamp(t time.Time) int64 {
	utcMilliseconds := t.UnixMilli()
	milliseconds := utcMilliseconds + utcToLocal*millisecondsPerSecond
	return milliseconds
}

// Time 本地毫秒数转time.Time
//
//	local 转 UTC
func Time(milliseconds int64) time.Time {
	utcMilliseconds := milliseconds + localToUTC*millisecondsPerSecond
	return time.UnixMilli(utcMilliseconds)
}
