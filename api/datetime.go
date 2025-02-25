package api

import (
	"errors"
	"strings"
	"time"
)

const (
	YearOnly = "2006" // 日期格式仅输出"年"
)

var (
	ErrDateFormat = errors.New("日期格式无法确定")
)

// ParseTime 解析时间字符串
func ParseTime(timestr string) (time.Time, error) {
	s := strings.TrimSpace(timestr)
	switch len(s) {
	case len(DateFormat):
		return time.ParseInLocation(DateFormat, s, time.Local)
	case len(DateFormat2):
		return time.ParseInLocation(DateFormat2, s, time.Local)
	case len(DateFormat3):
		return time.ParseInLocation(DateFormat3, s, time.Local)
	case len(TimeFormat2):
		return time.ParseInLocation(TimeFormat2, s, time.Local)
	case len(TimeFormat):
		return time.ParseInLocation(TimeFormat, s, time.Local)
	case len(Timestamp):
		return time.ParseInLocation(Timestamp, s, time.Local)
	default:
		return time.Time{}, ErrDateFormat
	}
}

// DifferDays 计算天数差
//
//	从 t1 回到 t2 需要多少天
func DifferDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

// DateZero t 的0点0分0秒
func DateZero(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// IsWorkday 是否工作日
func IsWorkday(t time.Time) bool {
	weekDay := t.Weekday()
	if weekDay == time.Sunday || weekDay == time.Saturday {
		return false
	}
	return true
}

// GetMonthDay 获得当前月的初始和结束日期
func GetMonthDay(date ...string) (string, string) {
	now := time.Now()
	if len(date) > 0 {
		day, err := ParseTime(date[0])
		if err == nil {
			now = day
		}
	}
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	f := firstOfMonth.Unix()
	l := lastOfMonth.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

// GetWeekDay 获得当前周的初始和结束日期
func GetWeekDay(date ...string) (string, string) {
	now := time.Now()
	if len(date) > 0 {
		day, err := ParseTime(date[0])
		if err == nil {
			now = day
		}
	}
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Sunday = 0
	if offset > 0 {
		offset = -6
	}

	lastoffset := int(time.Saturday - now.Weekday())
	//周日做特殊判断 因为time.Sunday = 0
	if lastoffset == 6 {
		lastoffset = -1
	}

	firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

// GetQuarterDay 获得当前季度的初始和结束日期
func GetQuarterDay(months ...int) (string, string) {
	mn := 0
	if len(months) > 0 {
		mn = months[0]
	}
	now := time.Now().AddDate(0, -mn, 0)
	year := now.Format(YearOnly)
	month := int(now.Month())
	var firstOfQuarter string
	var lastOfQuarter string
	if month >= 1 && month <= 3 {
		//1月1号
		firstOfQuarter = year + "-01-01 00:00:00"
		lastOfQuarter = year + "-03-31 23:59:59"
	} else if month >= 4 && month <= 6 {
		firstOfQuarter = year + "-04-01 00:00:00"
		lastOfQuarter = year + "-06-30 23:59:59"
	} else if month >= 7 && month <= 9 {
		firstOfQuarter = year + "-07-01 00:00:00"
		lastOfQuarter = year + "-09-30 23:59:59"
	} else {
		firstOfQuarter = year + "-10-01 00:00:00"
		lastOfQuarter = year + "-12-31 23:59:59"
	}
	return firstOfQuarter, lastOfQuarter
}

// GetQuarterDayByDate 通过给定的日期 获得日期所在上一个季度的初始和结束日期
//
//	diffQuarters 季度偏移数, 大于0前移diffQuarters个季度, 小于0后移diffQuarters个季度, 默认为当前季度
func GetQuarterDayByDate(date string, diffQuarters ...int) (firstOfQuarter, lastOfQuarter string) {
	diff := 1
	if len(diffQuarters) > 0 {
		diff = diffQuarters[0]
	}
	_, firstOfQuarter, lastOfQuarter = GetQuarterByDate(date, diff)
	return firstOfQuarter, lastOfQuarter
}

// GetQuarterByDate 通过给定的日期 获得日期所在财报的季度、初始以及结束日期
//
//	diffQuarters 季度偏移数, 大于0前移diffQuarters个季度, 小于0后移diffQuarters个季度, 默认为当前季度
func GetQuarterByDate(date string, diffQuarters ...int) (quarter, first, last string) {
	diff := 0
	if len(diffQuarters) > 0 {
		diff = diffQuarters[0]
	}
	now, _ := ParseTime(date)
	now = now.AddDate(0, -3*diff, 0)
	year := now.Format(YearOnly)
	month := int(now.Month())
	var firstOfQuarter string
	var lastOfQuarter string
	if month >= 1 && month <= 3 {
		//1月1号
		firstOfQuarter = year + "-01-01 00:00:00"
		lastOfQuarter = year + "-03-31 23:59:59"
		quarter = year + "Q1"
	} else if month >= 4 && month <= 6 {
		firstOfQuarter = year + "-04-01 00:00:00"
		lastOfQuarter = year + "-06-30 23:59:59"
		quarter = year + "Q2"
	} else if month >= 7 && month <= 9 {
		firstOfQuarter = year + "-07-01 00:00:00"
		lastOfQuarter = year + "-09-30 23:59:59"
		quarter = year + "Q3"
	} else {
		firstOfQuarter = year + "-10-01 00:00:00"
		lastOfQuarter = year + "-12-31 23:59:59"
		quarter = year + "Q4"
	}
	return quarter, firstOfQuarter, lastOfQuarter
}
