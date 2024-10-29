package livetest

import (
	"reflect"
	"testing"
)

func Test_possibleStringCount(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				word: "abbcccc",
			},
			want: 5,
		},

		{
			name: "2",
			args: args{
				word: "abc",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleStringCount(tt.args.word); got != tt.want {
				t.Errorf("possibleStringCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSubtreeSizes(t *testing.T) {
	type args struct {
		parent []int
		s      string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				parent: []int{-1, 0, 0, 1, 1, 1},
				s:      "abaabc",
			},
			want: []int{6, 3, 1, 1, 1, 1},
		},

		{
			name: "2",
			args: args{
				parent: []int{-1, 0, 4, 0, 1},
				s:      "abbba",
			},
			want: []int{5, 2, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubtreeSizes(tt.args.parent, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubtreeSizes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxScore(t *testing.T) {
	type args struct {
		n           int
		k           int
		stayScore   [][]int
		travelScore [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n:         2,
				k:         1,
				stayScore: [][]int{{2, 3}},
				travelScore: [][]int{
					{0, 2},
					{1, 0},
				},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				n: 3,
				k: 2,
				stayScore: [][]int{{3, 4, 2},
					{2, 1, 2}},
				travelScore: [][]int{
					{0, 1, 2},
					{2, 0, 4},
					{3, 2, 0},
				},
			},
			want: 8,
		},

		{
			name: "3",
			args: args{
				n:         1,
				k:         2,
				stayScore: [][]int{{1}, {2}},
				travelScore: [][]int{
					{0},
				},
			},
			want: 3,
		},

		{
			name: "4",
			args: args{
				n:         2,
				k:         1,
				stayScore: [][]int{{1, 1}},
				travelScore: [][]int{
					{0, 1}, {0, 6},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.n, tt.args.k, tt.args.stayScore, tt.args.travelScore); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findOriginalTypedString(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				word: "aabbccdd",
				k:    7,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOriginalTypedString(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("findOriginalTypedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
