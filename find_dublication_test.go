package livetest

import "testing"

func TestFindDuplication(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 3, 4, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("FindDuplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
