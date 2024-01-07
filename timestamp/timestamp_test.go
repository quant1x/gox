package timestamp

import (
	"fmt"
	"gitee.com/quant1x/gox/api"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	now := Now()
	fmt.Println("today", now-now%millisecondsPerDay)
	fmt.Println((now % millisecondsPerDay) / millisecondsPerHour)
	fmt.Println((now % millisecondsPerHour) / millisecondsPerMinute)
	fmt.Println((now % millisecondsPerMinute) / millisecondsPerSecond)
	fmt.Println(now % millisecondsPerSecond)

	fmt.Println(now % millisecondsPerDay)

	date := "2024-01-07"
	tm, err := api.ParseTime(date)
	fmt.Println(tm, err)
	fmt.Println(tm.Date())
	ts := tm.Local().Unix()
	ts = tm.UnixMilli()
	fmt.Println("today1 :", ts)
	ts = tm.Local().UnixMilli()
	fmt.Println("today2 :", ts+int64(offsetInSecondsEastOfUTC*millisecondsPerSecond))
	t1 := time.Unix(ts/1000, 0)
	fmt.Println("t1 =>", t1)
	t3 := Timestamp(tm)
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
