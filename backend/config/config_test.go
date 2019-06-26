package config

import "testing"

func Test_randomXid_Do(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		x       *randomXid
		args    args
		wantLen int
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			args:    args{length: 10},
			wantLen: 10,
		},
		{
			name:    "success over 20(xid generated string 20)",
			args:    args{length: 21},
			wantLen: 21,
		},
		{
			name:    "zero",
			args:    args{length: 0},
			wantLen: 0,
		},
		{
			name:    "minus",
			args:    args{length: -1},
			wantLen: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &randomXid{}
			if got := x.Do(tt.args.length); len(got) != tt.wantLen {
				t.Errorf("not len equal randomXid.Do() = %v, want %v", got, tt.wantLen)
			}
		})
	}
}
