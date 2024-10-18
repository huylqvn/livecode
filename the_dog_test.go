package livetest

import "testing"

func TestDogFood(t *testing.T) {
	type args struct {
		strArr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				strArr: []string{"OOOO", "OOFF", "OCHO", "OFOO"},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DogFood(tt.args.strArr); got != tt.want {
				t.Errorf("DogFood() = %v, want %v", got, tt.want)
			}
		})
	}
}
