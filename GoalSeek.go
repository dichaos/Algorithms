package main

import (
	"MathD/MathHelpers"
	"math"
)

func Newton(seekValue float64, maxIterations int, a float64, cutOff float64, f MathHelpers.Fn, derF MathHelpers.Fn) float64 {
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

func Bisectional(TOL float64, maxIterations int, a float64, b float64, seekValue float64, f MathHelpers.Fn) float64 {
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
