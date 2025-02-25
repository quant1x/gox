package api

import (
	"fmt"
	"time"
)

// GetWeekRange 获取指定日期所在周的起止时间范围（周一到周日）
// 参数 date 为可选参数（最多一个），格式为 "2006-01-02"，缺省使用当前时间
// 返回值格式为 "2006-01-02 00:00:00", "2006-01-02 23:59:59"
func GetWeekRange(date ...string) (startStr, endStr string, err error) {
	// 参数合法性检查
	switch {
	case len(date) > 1:
		return "", "", fmt.Errorf("参数数量错误：最多接受1个日期参数")
	case len(date) == 1 && len(date[0]) != 10: // 快速失败检查
		return "", "", fmt.Errorf("日期格式错误：需要10位字符")
	}

	// 时区预处理（显式获取本地时区）
	localLoc := time.Local
	now := time.Now().In(localLoc)

	// 基准时间初始化
	var baseTime time.Time
	if len(date) > 0 && date[0] != "" {
		if baseTime, err = time.ParseInLocation(time.DateOnly, date[0], localLoc); err != nil {
			return "", "", fmt.Errorf("日期解析失败：%w", err)
		}
		baseTime = baseTime.In(localLoc) // 确保时区一致性
	} else {
		baseTime = now
	}

	// 核心计算逻辑（封装为闭包提升可读性）
	computeRange := func(t time.Time) (start, end time.Time) {
		// 计算周一零点
		weekdayOffset := int(t.Weekday()-time.Monday+7) % 7
		monday := t.AddDate(0, 0, -weekdayOffset)

		start = time.Date(
			monday.Year(), monday.Month(), monday.Day(),
			0, 0, 0, 0,
			localLoc,
		)

		// 计算周日结束
		sunday := start.AddDate(0, 0, 6)
		end = time.Date(
			sunday.Year(), sunday.Month(), sunday.Day(),
			23, 59, 59, 0,
			localLoc,
		)
		return
	}

	// 执行计算并格式化
	start, end := computeRange(baseTime)
	return start.Format(time.DateTime), end.Format(time.DateTime), nil
}
