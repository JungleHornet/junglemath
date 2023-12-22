package junglemath

import (
	"math"
	"strconv"
)

func Pythag(leg1, leg2 float64) string {

	leg1 = leg1 * leg1
	leg2 = leg2 * leg2

	hyp := math.Sqrt(leg1 + leg2)

	return strconv.FormatFloat(hyp, 'f', -1, 64)
}
