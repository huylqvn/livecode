package livetest

import "testing"

func Test_movesSolution(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "^>^^<>",
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				s: "<<<<",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movesSolution(tt.args.s); got != tt.want {
				t.Errorf("movesSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
