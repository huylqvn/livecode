package livetest

import "testing"

func TestFixHole(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				a: []int{11, 20, 9, 15},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FixHole(tt.args.a)
			t.Log("res", got)
		})
	}
}
