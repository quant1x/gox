// Package timestamp 本地时间戳相关功能
package timestamp

import (
	"fmt"
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

var (
	defaultLocal = time.Local
	// 获取偏移的秒数
	zoneName, offsetInSecondsEastOfUTC = time.Now().Zone()
	_                                  = zoneName
	// 本地0秒
	zeroTime = time.Date(1970, 1, 1, 0, 0, 0, 0, defaultLocal)
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
	return v2Now()
}

func v1Now() int64 {
	sec, nsec, _ := now()
	sec += int64(offsetInSecondsEastOfUTC)
	milli := sec*MillisecondsPerSecond + int64(nsec)/1e6%MillisecondsPerSecond
	return milli
}

// TimeToTimestamp 获取time.Time的本地毫秒数
//
//	UTC 转 local
func TimeToTimestamp(t time.Time) int64 {
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
	milliseconds := TimeToTimestamp(t)
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

const (
	secondsPerMinute = 60
	secondsPerHour   = 60 * secondsPerMinute
	secondsPerDay    = 24 * secondsPerHour
	secondsPerWeek   = 7 * secondsPerDay
	daysPer400Years  = 365*400 + 97
	daysPer100Years  = 365*100 + 24
	daysPer4Years    = 365*4 + 1
)

const (
	// The unsigned zero year for internal calculations.
	// Must be 1 mod 400, and times before it will not compute correctly,
	// but otherwise can be changed at will.
	absoluteZeroYear = -292277022399

	// The year of the zero Time.
	// Assumed by the unixToInternal computation below.
	internalYear = 1

	// Offsets to convert between internal and absolute or Unix times.
	absoluteToInternal int64 = (absoluteZeroYear - internalYear) * 365.2425 * secondsPerDay
	internalToAbsolute       = -absoluteToInternal

	unixToInternal int64 = (1969*365 + 1969/4 - 1969/100 + 1969/400) * secondsPerDay
	internalToUnix int64 = -unixToInternal

	wallToInternal int64 = (1884*365 + 1884/4 - 1884/100 + 1884/400) * secondsPerDay
)

// Timestamp 时间戳类型
type Timestamp int64

// abs returns the time t as an absolute time, adjusted by the zone offset.
// It is called when computing a presentation property like Month or Hour.
func (t Timestamp) abs() uint64 {
	//l := defaultLocal
	ms := int64(t / MillisecondsPerSecond)
	return uint64(ms + (unixToInternal + internalToAbsolute))
}

// 调用公开结构的私有方法
//
//go:linkname absDate time.absDate
func absDate(abs uint64, full bool) (year int, month time.Month, day int, yday int)

// DateTime 获取日期时间毫秒
func (t Timestamp) DateTime() (year, month, day, hour, minute, second, millisecond int) {
	ms := int64(t)
	absSeconds := t.abs()
	year, m, day, yday := absDate(absSeconds, true)
	_ = yday
	month = int(m)
	hour = int((ms % MillisecondsPerDay) / MillisecondsPerHour)
	minute = int((ms % MillisecondsPerHour) / MillisecondsPerMinute)
	second = int((ms % MillisecondsPerMinute) / MillisecondsPerSecond)
	millisecond = int(ms % MillisecondsPerSecond)
	return
}

func (t Timestamp) String() string {
	year, month, day, hour, minute, second, millisecond := t.DateTime()
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%03d", year, month, day, hour, minute, second, millisecond)
}
