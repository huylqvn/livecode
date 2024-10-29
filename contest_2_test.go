package livetest

import (
	"testing"
)

func Test_maxFactorScore(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{
				nums: []int{11, 4},
			},
			want: 121,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFactorScore(tt.args.nums); got != tt.want {
				t.Errorf("maxFactorScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthAfterTransformations(t *testing.T) {
	type args struct {
		s string
		t int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "jqktcurgdvlibczdsvnsg",
				t: 7517,
			},
			want: 79033769,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthAfterTransformations(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("lengthAfterTransformations() = %v, want %v", got, tt.want)
			}
		})
	}
}
