package junglemath

import (
	"math"
	"strconv"
)

func Pythag(leg1, leg2 float64, format format) string {

	leg1 = leg1 * leg1
	leg2 = leg2 * leg2

	hyp := math.Sqrt(leg1 + leg2)

	if format == Decimal {
		return strconv.FormatFloat(hyp, 'f', -1, 64)
	}

	root := math.Round(hyp * hyp)
	sqrtHyp := "√" + strconv.FormatFloat(root, 'f', -1, 64)

	if format == Radical {
		return sqrtHyp
	}

	rootCoefficient := int64(1)
	simpleRootInt := root

	for i := float64(2); i <= math.Round(math.Sqrt(root)); i++ {
		if (simpleRootInt / i) == math.Trunc(simpleRootInt/i) {
			if !(i == simpleRootInt) && !(i == 1) {
				if math.Sqrt(simpleRootInt/i) == math.Trunc(math.Sqrt(simpleRootInt/i)) {
					simpleRootInt = i
					rootCoefficient = rootCoefficient * int64(math.Sqrt(root/i))
				}
			}
		}
	}

	simpleRoot := strconv.FormatInt(rootCoefficient, 10) + "√" +
		strconv.FormatFloat(simpleRootInt, 'f', -1, 64)

	if format == SimplifiedRadical {
		return simpleRoot
	}

	return "Error: format is not dec, rad, or simpRad"

}
