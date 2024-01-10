// Package timestamp 本地时间戳相关功能
package timestamp

import (
	"time"
	_ "unsafe" // for go:linkname
)

const (
	SecondsPerMinute      = 60
	SecondsPerHour        = 60 * SecondsPerMinute
	SecondsPerDay         = 24 * SecondsPerHour
	MillisecondsPerSecond = 1000
	MillisecondsPerMinute = SecondsPerMinute * MillisecondsPerSecond
	MillisecondsPerHour   = SecondsPerHour * MillisecondsPerSecond
	MillisecondsPerDay    = SecondsPerDay * MillisecondsPerSecond
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
	// 本地0秒
	zeroTime = time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)
	// UTC到本地的偏移秒数
	utcToLocal = int64(offsetInSecondsEastOfUTC)
	// 本地到UTC的偏移秒数
	localToUTC = -utcToLocal
	// 偏移的毫秒数
	//offsetMilliseconds = offsetInSecondsEastOfUTC * MillisecondsPerSecond
)

// Now 获取本地当前的时间戳, 毫秒数
//
//	UTC 转 local
func Now() int64 {
	sec, nsec, _ := now()
	sec += int64(offsetInSecondsEastOfUTC)
	milli := sec*MillisecondsPerSecond + int64(nsec)/1e6%MillisecondsPerSecond
	return milli
}

// Timestamp 获取time.Time的本地毫秒数
//
//	UTC 转 local
func Timestamp(t time.Time) int64 {
	utcMilliseconds := t.UnixMilli()
	milliseconds := utcMilliseconds + utcToLocal*MillisecondsPerSecond
	return milliseconds
}

// Time 本地毫秒数转time.Time
//
//	local 转 UTC
func Time(milliseconds int64) time.Time {
	utcMilliseconds := milliseconds + localToUTC*MillisecondsPerSecond
	return time.UnixMilli(utcMilliseconds)
}

// SinceZero 从0点到milliseconds过去的毫秒数
func SinceZero(milliseconds int64) int64 {
	elapsed := milliseconds % MillisecondsPerDay
	if elapsed < 0 {
		elapsed += MillisecondsPerDay
	}
	return elapsed
}

// ZeroHour 零点整的毫秒数
func ZeroHour(milliseconds int64) int64 {
	elapsed := SinceZero(milliseconds)
	diff := milliseconds - elapsed
	return diff
}

// Since t当日0点整到t的毫秒数
func Since(t time.Time) int64 {
	milliseconds := Timestamp(t)
	elapsed := SinceZero(milliseconds)
	return elapsed
}

// Today 获取当日0点整的时间戳, 毫秒数
//
//	UTC 转 local
func Today() int64 {
	milliseconds := Now()
	//elapsed := milliseconds - milliseconds%MillisecondsPerDay
	elapsed := ZeroHour(milliseconds)
	return elapsed
}

// CurrentDateZero t日期的0点整
func CurrentDateZero(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

// TodayZero 这也是一个当日0点的用法
func TodayZero() time.Time {
	now := time.Now()
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

// SinceZeroHour t当天0点开始到t时的毫秒数
func SinceZeroHour(t time.Time) int64 {
	zero := CurrentDateZero(t)
	return t.Sub(zero).Milliseconds()
}
