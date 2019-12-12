package utils

import (
	"testing"
	"time"
)

func TestStrTime(t *testing.T) {
	type args struct {
		atime int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				atime: time.Now().Unix() - 10,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrTime(tt.args.atime); got != tt.want {
				t.Errorf("StrTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeString(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				args: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeString(tt.args.args...); got != tt.want {
				t.Errorf("mergeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
