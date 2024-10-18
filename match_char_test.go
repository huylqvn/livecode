package livetest

import "testing"

func TestMatchingCharacters(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				str: "abccdefghi",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchingCharacters(tt.args.str); got != tt.want {
				t.Errorf("MatchingCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
