package api

import "testing"

func TestParseFloat(t *testing.T) {
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
			if got := ParseFloat(tt.args.s); got != tt.want {
				t.Errorf("ParseFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseUint(t *testing.T) {
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
			if got := ParseUint(tt.args.s); got != tt.want {
				t.Errorf("ParseUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseint(t *testing.T) {
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
			if got := ParseInt(tt.args.s); got != tt.want {
				t.Errorf("ParseInt() = %v, want %v", got, tt.want)
			}
		})
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
