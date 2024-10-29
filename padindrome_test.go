package livetest

import "testing"

func TestPalindromeSolution(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "a?a",
				k: 1,
			},
			want: "aaa",
		},
		{
			name: "2",
			args: args{
				s: "????",
				k: 2,
			},
			want: "aaaa",
		},
		{
			name: "3",
			args: args{
				s: "abc?abcaa",
				k: 3,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PalindromeSolution(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("PalindromeSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
