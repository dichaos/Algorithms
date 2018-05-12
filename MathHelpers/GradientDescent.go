package MathHelpers

import (
	"math"
	"math/rand"
)

func GradientDescendent(stepSize float64, precision float64, dfx Fn) float64 {
	x := rand.Float64()
	previous_step_size := 1 / precision
	var prev_x float64

	for previous_step_size > precision {
		prev_x = x
		x += -stepSize * dfx(prev_x)
		previous_step_size = math.Abs(x - prev_x)
	}

	return x
}
