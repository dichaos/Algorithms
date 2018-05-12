package main

import (
	"MathD/MathHelpers"
	"fmt"
	"math"
)

func main() {

	fx := MathHelpers.Fn(func(x float64) float64 {
		return math.Pow(x, 3) - x - 2
	})

	dfx := MathHelpers.Fn(func(x float64) float64 {
		return 3*math.Pow(x, 2) - 1
	})

	res := Bisectional(0.00001, 1000, 1, 2, 0, fx)
	resNewton := Newton(0, 1000, 1, 0.00001, fx, dfx)

	fmt.Printf("Bisectional: %v,%v", res, fx(res))
	fmt.Println()
	fmt.Printf("Newton: %v,%v", resNewton, fx(resNewton))
	fmt.Println()
	fmt.Printf("GradientDescent: %v,%v", resNewton, fx(resNewton))
	fmt.Println()

	var x [][]float64
	MathHelpers.DistanceMatrix(x, MathHelpers.ManhatanDistance)
}
