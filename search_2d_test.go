package livetest

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "1",
		// 	args: args{
		// 		matrix: [][]int{
		// 			{1, 3, 5, 7},
		// 			{10, 11, 16, 20},
		// 			{23, 30, 34, 50},
		// 		},
		// 		target: 3,
		// 	},
		// 	want: true,
		// },
		{
			name: "2",
			args: args{
				matrix: [][]int{
					{-5},
				},
				target: -5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
