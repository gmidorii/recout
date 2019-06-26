package app

import (
	"testing"
	"time"
)

func Test_subDate(t *testing.T) {
	type args struct {
		before time.Time
		after  time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "pass 0 day.",
			args: args{
				before: time.Date(2019, time.May, 9, 0, 0, 0, 0, time.UTC),
				after:  time.Date(2019, time.May, 9, 23, 59, 0, 0, time.UTC),
			},
			want: 0,
		},
		{
			name: "pass 1 day.",
			args: args{
				before: time.Date(2019, time.May, 9, 0, 0, 0, 0, time.UTC),
				after:  time.Date(2019, time.May, 10, 0, 0, 0, 0, time.UTC),
			},
			want: 1,
		},
		{
			name: "pass 1 month.",
			args: args{
				before: time.Date(2019, time.May, 9, 0, 0, 0, 0, time.UTC),
				after:  time.Date(2019, time.June, 9, 0, 0, 0, 0, time.UTC),
			},
			want: 31,
		},
		{
			name: "past day.",
			args: args{
				before: time.Date(2019, time.May, 9, 0, 0, 0, 0, time.UTC),
				after:  time.Date(2019, time.May, 8, 23, 59, 0, 0, time.UTC),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subDate(tt.args.before, tt.args.after); got != tt.want {
				t.Errorf("subDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
