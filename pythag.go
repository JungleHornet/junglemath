package junglemath

import (
	"math"
)

func Pythag(leg1, leg2 float64) float64 {

	leg1 = leg1 * leg1
	leg2 = leg2 * leg2

	hyp := math.Sqrt(leg1 + leg2)

	return hyp
}
