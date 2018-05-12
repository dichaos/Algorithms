package MathHelpers

import "math/rand"

type Shuffle func(x []int) []int

func FisherYates(x []int) []int {
	i := 0

	for range x {
		pos := rand.Intn(len(x))
		x[i], x[pos] = x[pos], x[i]
		i++
	}

	return x
}

func SattoloCycle(x []int) []int {
	for i := len(x) - 1; i > 0; i-- {
		j := rand.Intn(i)
		x[i], x[j] = x[j], x[i]
	}

	return x
}
