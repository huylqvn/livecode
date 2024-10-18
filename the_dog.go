package livetest

import (
	"math"
)

func Point(x, y int) []int {
	return []int{x, y}
}

type DogState struct {
	xCur  int
	yCur  int
	xDog  int
	yDog  int
	Map2D [][]int
}

func calculateDistance(x1, y1, x2, y2 int) int {
	if x1 == x2 || y1 == y2 {
		return int(math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))
	}

	xDistance := int(math.Abs(float64(y2 - y1)))
	yDistance := int(math.Abs(float64(x2 - x1)))
	return xDistance + yDistance
}

func create2DMap(strArr []string) *DogState {
	dogState := &DogState{}

	dogState.Map2D = [][]int{}

	for row := 0; row < len(strArr); row++ {
		for col := 0; col < len(strArr[row]); col++ {
			if strArr[row][col] == 'H' {
				dogState.xCur = row
				dogState.yCur = col
			} else if strArr[row][col] == 'C' {
				dogState.xDog = row
				dogState.yDog = col
			} else if strArr[row][col] == 'F' {
				dogState.Map2D = append(dogState.Map2D, Point(row, col))
			}
		}
	}

	return dogState
}

func DogFood(strArr []string) int {
	dog := create2DMap(strArr)

	totalDistance := 0
	for len(dog.Map2D) > 0 {
		min := math.MaxInt32
		var index, minx, miny int

		for i := 0; i < len(dog.Map2D); i++ {
			distance := calculateDistance(dog.xCur, dog.yCur, dog.Map2D[i][0], dog.Map2D[i][1])
			if distance < min {
				min = distance
				index = i
				minx = dog.Map2D[i][0]
				miny = dog.Map2D[i][1]
			}
			if distance == min && calculateDistance(dog.xDog, dog.yDog, minx, miny) < calculateDistance(dog.xDog, dog.yDog, dog.Map2D[i][0], dog.Map2D[i][1]) {
				min = distance
				index = i
				minx = dog.Map2D[i][0]
				miny = dog.Map2D[i][1]
			}
		}

		totalDistance += min

		dog.xCur = minx
		dog.yCur = miny
		dog.Map2D = append(dog.Map2D[:index], dog.Map2D[index+1:]...)
	}

	totalDistance += calculateDistance(dog.xCur, dog.yCur, dog.xDog, dog.yDog)
	return totalDistance
}
