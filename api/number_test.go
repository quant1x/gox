package api

import "testing"

func TestParseFloat(t *testing.T) {
	s := "123.456"
	var f64 float64
	f64 = ParseFloat(s)
	if f64 != 123.456 {
		t.Errorf("解析float64到int64失败")
	}

	s = "123.00"
	f64 = ParseFloat(s)
	if f64 != 123.00 {
		t.Errorf("解析float64到int64失败")
	}
}

func TestParseUint(t *testing.T) {
	s := "123.456"
	var i64 uint64
	i64 = ParseUint(s)
	if i64 != 123 {
		t.Errorf("解析float64到int64失败")
	}
}

func TestParseint(t *testing.T) {
	s := "123.456"
	var i64 int64
	i64 = ParseInt(s)
	if i64 != 123 {
		t.Errorf("解析float64到int64失败")
	}
}

func Test_parseBestEffort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBestEffort(tt.args.s); got != tt.want {
				t.Errorf("parseBestEffort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInt64BestEffort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInt64BestEffort(tt.args.s); got != tt.want {
				t.Errorf("parseInt64BestEffort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseUint64BestEffort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseUint64BestEffort(tt.args.s); got != tt.want {
				t.Errorf("parseUint64BestEffort() = %v, want %v", got, tt.want)
			}
		})
	}
}
