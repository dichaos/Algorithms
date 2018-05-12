package MathHelpers

import "math"

type Distance func(x []float64, y []float64) float64

func ManhatanDistance(x []float64, y []float64) float64 {
	var distance float64
	i := 0

	for range x {
		distance += math.Abs(x[i] - y[i])
		i++
	}

	return distance
}

func EucledianDistance(x []float64, y []float64) float64 {
	var distance float64
	i := 0

	for range x {
		distance += math.Pow(math.Abs(x[i]-y[i]), 2)
		i++
	}

	return math.Sqrt(distance)
}

func DistanceMatrix(arr [][]float64, distanceMethod Distance) [][]float64 {
	x := 0
	y := 0

	var toReturn [][]float64
	for range arr {
		for range arr[x] {
			if x == y {
				toReturn[x][y] = 0
			} else {
				toReturn[x][y] = distanceMethod(arr[x], arr[y])
			}
			y++
		}
		x++
	}

	return toReturn
}
