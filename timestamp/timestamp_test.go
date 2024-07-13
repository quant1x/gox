package timestamp

import (
	"fmt"
	"gitee.com/quant1x/gox/api"
	//"github.com/golang-module/carbon/v2"
	"testing"
	"time"
)

func TestTimePrivate(t *testing.T) {
	v := unixToInternal
	fmt.Println(v)
}

func TestYMDHMS_SSS(t *testing.T) {
	ms := Now()
	ts := Timestamp(ms)
	fmt.Println(ts.DateTime())
	fmt.Println(ts)
}

func TestNow(t *testing.T) {
	now := Now()
	fmt.Println("today1 :", now-now%MillisecondsPerDay)
	fmt.Println("today2 :", ZeroHour(now))
	fmt.Println("today3 :", Today())
	fmt.Println("  h :", (now%MillisecondsPerDay)/MillisecondsPerHour)
	fmt.Println("  m :", (now%MillisecondsPerHour)/MillisecondsPerMinute)
	fmt.Println("  s :", (now%MillisecondsPerMinute)/MillisecondsPerSecond)
	fmt.Println("sss :", now%MillisecondsPerSecond)

	fmt.Println(now % MillisecondsPerDay)

	date := "2024-01-07"
	tm, err := api.ParseTime(date)
	fmt.Println(tm, err)
	fmt.Println(tm.Date())
	ts := tm.Local().Unix()
	ts = tm.UnixMilli()
	fmt.Println("today1 :", ts)
	ts = tm.Local().UnixMilli()
	fmt.Println("today2 :", ts+int64(offsetInSecondsEastOfUTC*MillisecondsPerSecond))
	t1 := time.Unix(ts/1000, 0)
	fmt.Println("t1 =>", t1)
	t3 := TimeToTimestamp(tm)
	fmt.Println("t3 =>", t3)
	t4 := Time(t3)
	fmt.Println("t4 =>", t4)

	t2 := time.UnixMilli(now) //.UTC()
	fmt.Println("t2 =>", t2)
}

func TestTime1(t *testing.T) {
	formatTimeStr := "2017-04-11 13:33:37"
	formatTime, err := time.Parse("2006-01-02 15:04:05", formatTimeStr)
	if err == nil {
		fmt.Println(formatTime) //打印结果：2017-04-11 13:33:37 +0000 UTC
		fmt.Println(formatTime.Unix())
	}
}

//func TestTime2(t *testing.T) {
//	// 今天此刻
//	fmt.Printf("%s", carbon.Now()) // 2020-08-05 13:14:15
//}

func TestTime3(t *testing.T) {
	t1 := zeroTime.UnixMilli()
	fmt.Println(t1)
	t2 := time.Now()
	fmt.Println(t2)
	t3 := time.Since(t2)
	fmt.Println(t3.Seconds())
}

func TestSinceZeroHour(t *testing.T) {
	now := time.Now()
	t1 := SinceZeroHour(now)
	fmt.Println(t1)
	t2 := Since(now)
	fmt.Println(t2)
}

func BenchmarkTimestamp_release(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tm := Now()
		_ = tm
	}
}

func BenchmarkTimestamp_v0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		now := time.Now()
		tm := now.UnixMilli()
		_ = tm
	}
}

func BenchmarkTimestamp_v1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tm := v1Now()
		_ = tm
	}
}

func BenchmarkTimestamp_v2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tm := v2Now()
		_ = tm
	}
}

func add_two(n int32) int32 {
	return n + 2
}

func BenchmarkTimestamp_add_two(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add_two(int32(i))
	}
}
