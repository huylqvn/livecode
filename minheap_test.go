package livetest

import "testing"

func TestHandleMinHeap(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: []int{11, 20, 9, 15},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleMinHeap(tt.args.a); got != tt.want {
				t.Errorf("HandleMinHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
