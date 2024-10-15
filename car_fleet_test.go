package livetest

import "testing"

func TestCarFleet(t *testing.T) {
	type args struct {
		target   int
		position []int
		speed    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				target:   20,
				position: []int{6, 2, 17},
				speed:    []int{3, 9, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CarFleet(tt.args.target, tt.args.position, tt.args.speed); got != tt.want {
				t.Errorf("CarFleet() = %v, want %v", got, tt.want)
			}
		})
	}
}
