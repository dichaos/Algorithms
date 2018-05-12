package MultidimensionalScaling

import (
	"MathD/MathHelpers"
	"math"
)

func SammonsMapping(distances [][]float64, outputDimension int, maxIteration int, distance MathHelpers.Distance, shuffle MathHelpers.Shuffle) [][]float64 {
	var lambda float64 = 1
	toReturn := make([][]float64, len(distances))
	count := 0

	for count < maxIteration {
		indicesx := makeSuffleRange(0, len(distances), shuffle)
		indicesy := makeSuffleRange(0, len(distances), shuffle)

		x := 0
		y := 0
		for range indicesx {

			for range indicesy {
				if indicesx[x] == indicesy[y] {
					continue
				}

				dij := distances[indicesx[x]][indicesy[y]]
				Dij := distance(toReturn[indicesx[x]], toReturn[indicesy[y]])

				if Dij == 0 {
					Dij = 1e-10
				}

				delta := lambda * (dij - Dij) / Dij

				k := 0
				for range toReturn[indicesy[y]] {

					correction := delta * (toReturn[indicesx[x]][k] - toReturn[indicesx[y]][k])

					toReturn[indicesx[x]][k] += correction
					toReturn[indicesx[y]][k] -= correction

					k++
				}

				y++
			}
			x++
		}

		count++
		reduceLambda(&lambda, count, maxIteration)
	}

	return toReturn
}

func reduceLambda(lambda *float64, iteration int, maxIteration int) {
	ratio := float64(iteration) / float64(maxIteration)

	*lambda = math.Pow(0.01, ratio)
}

func makeSuffleRange(min, max int, shuffle MathHelpers.Shuffle) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return shuffle(a)
}
