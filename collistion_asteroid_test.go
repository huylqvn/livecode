package livetest

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				asteroids: []int{1, -2, -2, -2},
			},
			want: []int{-2, -2, -2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
