package livetest

import "testing"

func Test_revertString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				str: "hello world",
			},
			want: "olleh dlrow",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := revertString(tt.args.str); got != tt.want {
				t.Errorf("revertString() = %v, want %v", got, tt.want)
			}
		})
	}
}
