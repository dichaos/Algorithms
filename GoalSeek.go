package main

import (
	"fmt"
	"math"
)

type fn func(float64) float64

func main() {

	function := fn(func(x float64) float64 {
		return math.Pow(x, 3) - x - 2
	})

	derivative := fn(func(x float64) float64 {
		return 3*math.Pow(x, 2) - 1
	})

	res := Bisectional(0.00001, 1000, 1, 2, 0, function)
	resNewton := Newton(0, 1000, 1, 0.00001, function, derivative)

	fmt.Printf("%v,%v", res, function(res))
	fmt.Println()
	fmt.Printf("%v,%v", resNewton, function(resNewton))

}

func Newton(seekValue float64, maxIterations int, a float64, cutOff float64, f fn, derF fn) float64 {
	x := a
	newX := x
	count := 0

	for math.Abs(seekValue-x) > cutOff && count < maxIterations {
		newX = x - (f(x) / derF(x))
		x = newX
		count++
	}

	return x
}

func Bisectional(TOL float64, maxIterations int, a float64, b float64, seekValue float64, f fn) float64 {
	for i := 0; i < maxIterations; i++ {
		c := (a + b) / 2
		fc := f(c)
		fa := f(a)

		if fc == seekValue || (b-a)/2 < TOL {
			return c
		}

		if (fc < seekValue && fa < seekValue) || (fc > seekValue && fa > seekValue) {
			a = c
		} else {
			b = c
		}
	}

	return a
}
