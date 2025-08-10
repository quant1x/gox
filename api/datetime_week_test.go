package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeekRange(t *testing.T) {
	type args struct {
		dateStrs []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "w0",
			args:    args{dateStrs: []string{}},
			want:    "2025-02-24 00:00:00",
			want1:   "2025-03-02 23:59:59",
			wantErr: nil,
		},
		{
			name:    "w1",
			args:    args{dateStrs: []string{"2025-02-24"}},
			want:    "2025-02-24 00:00:00",
			want1:   "2025-03-02 23:59:59",
			wantErr: nil,
		},
		{
			name:    "w2",
			args:    args{dateStrs: []string{"2025-02-25"}},
			want:    "2025-02-24 00:00:00",
			want1:   "2025-03-02 23:59:59",
			wantErr: nil,
		},
		{
			name:    "w3",
			args:    args{dateStrs: []string{"2025-02-26"}},
			want:    "2025-02-24 00:00:00",
			want1:   "2025-03-02 23:59:59",
			wantErr: nil,
		},
		{
			name:    "w4",
			args:    args{dateStrs: []string{"2025-02-27"}},
			want:    "2025-02-24 00:00:00",
			want1:   "2025-03-02 23:59:59",
			wantErr: nil,
		},
		{
			name:    "w5",
			args:    args{dateStrs: []string{"2025-02-28"}},
			want:    "2025-02-24 00:00:00",
			want1:   "2025-03-02 23:59:59",
			wantErr: nil,
		},
		{
			name:    "w6",
			args:    args{dateStrs: []string{"2025-03-01"}},
			want:    "2025-02-24 00:00:00",
			want1:   "2025-03-02 23:59:59",
			wantErr: nil,
		},
		{
			name:    "w7",
			args:    args{dateStrs: []string{"2025-03-02"}},
			want:    "2025-02-24 00:00:00",
			want1:   "2025-03-02 23:59:59",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, _ := GetWeekRange(tt.args.dateStrs...)
			//if !tt.wantErr(t, err, fmt.Sprintf("GetWeekRange(%v)", tt.args.dateStrs)) {
			//	return
			//}
			assert.Equalf(t, tt.want, got, "GetWeekRange(%v)", tt.args.dateStrs)
			assert.Equalf(t, tt.want1, got1, "GetWeekRange(%v)", tt.args.dateStrs)
		})
	}
}
