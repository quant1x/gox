package api

import (
	_ "cmp"
	"fmt"
	"math"
	"testing"
)

func TestDecimal0(t *testing.T) {
	f := -0.116
	fmt.Println(math.Signbit(f))
	fmt.Println(f)
	fmt.Printf("%.2f\n", f)
}

func TestDecimal(t *testing.T) {
	type args struct {
		value  float64
		digits []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "T9.8",
			args: args{
				value:  9.825,
				digits: []int{0},
			},
			want: 10,
		},
		{
			name: "T9.825",
			args: args{
				value:  9.825,
				digits: []int{2},
			},
			want: 9.83,
		},
		{
			name: "T0.116",
			args: args{
				value:  0.116,
				digits: []int{2},
			},
			want: 0.12,
		},
		{
			name: "T0.11",
			args: args{
				value:  0.1115355659035776,
				digits: []int{2},
			},
			want: 0.11,
		},
		{
			name: "T-0.11",
			args: args{
				value:  -0.1115355659035776,
				digits: []int{2},
			},
			want: -0.11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decimal(tt.args.value, tt.args.digits...); got != tt.want {
				t.Errorf("Decimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
