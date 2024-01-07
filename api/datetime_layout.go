package api

import "time"

const (
	// DateFormat  日期格式 yyyy-MM-dd
	DateFormat = "2006-01-02"
	// DateFormat2 日期格式 yyyyMMdd
	DateFormat2 = "20060102"
	// DateFormat3 日期格式 yyMMdd
	DateFormat3 = "060102"

	// TimeFormat 时间格式 yyyy-MM-dd HH:mm:ss
	TimeFormat = "2006-01-02 15:04:05"
	// TimeFormat2 时间格式 yyyyMMddHHmmss
	TimeFormat2 = "20060102150405"
	// Timestamp 时间戳 - 毫秒 时间格式 yyyy-MM-dd HH:mm:ss.SSS
	Timestamp = "2006-01-02 15:04:05.000"

	// TimeOnly 时分秒
	TimeOnly = time.TimeOnly
	// TimeAndMillisecond 时分秒毫
	TimeAndMillisecond = "15:04:05.000"

	PathDate = "20060102"   // 路径中的日期格式
	TextDate = "2006-01-02" // 数据中的日期格式
)
