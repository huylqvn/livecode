package livetest

import "testing"

func TestArrayJumping(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				arr: []int{1, 2, 3, 4, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayJumping(tt.args.arr); got != tt.want {
				t.Errorf("ArrayJumping() = %v, want %v", got, tt.want)
			}
		})
	}
}
