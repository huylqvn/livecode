package livetest

import "sort"

type Car struct {
	Pos, Speed int
	Time       float64
}

func CarFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 1 {
		return 1
	}

	res := 0
	cars := make([]Car, n)
	for i := 0; i < n; i++ {
		cars[i] = Car{
			Pos:   position[i],
			Speed: speed[i],
			Time:  float64(target-position[i]) / float64(speed[i]),
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Pos > cars[j].Pos
	})

	var max float64 = -1
	for i := 0; i < n; i++ {
		if cars[i].Time > max {
			res++
			max = cars[i].Time
		}
	}

	return res
}
