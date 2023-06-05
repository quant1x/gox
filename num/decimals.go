package num

import "math"

// Decimal 保留小数点四舍五入
func Decimal(value float64, digits ...int) float64 {
	defaultDigits := 2
	if len(digits) > 0 {
		defaultDigits = digits[0]
	}
	n10 := math.Pow10(defaultDigits)
	return math.Trunc((value+0.5/n10)*n10) / n10
}

// NetChangeRate 净增长率, 去掉百分比
func NetChangeRate[B Number, C Number](baseValue B, currentValue C) (changeRate float64) {
	changeRate = ChangeRate(baseValue, currentValue)
	changeRate = (changeRate - 1.00) * 100.00
	return changeRate
}

// ChangeRate 增长率
func ChangeRate[B Number, C Number](baseValue B, currentValue C) (changeRate float64) {
	return float64(currentValue) / float64(baseValue)
}
