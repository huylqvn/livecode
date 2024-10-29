package livetest

import "testing"

func TestPathHold(t *testing.T) {
	type args struct {
		l1 string
		l2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				l1: "..XX.X.",
				l2: "X.X.X..",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathHold(tt.args.l1, tt.args.l2); got != tt.want {
				t.Errorf("PathHold() = %v, want %v", got, tt.want)
			}
		})
	}
}
