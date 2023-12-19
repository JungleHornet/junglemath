package junglemath

import (
	"math"
	"strconv"
)

type format string

const (
	decimal           format = "dec"
	radical           format = "rad"
	simplifiedRadical format = "simpRad"
)

func CalcDistance(x1 float64, y1 float64, x2 float64, y2 float64, format format) string {

	/*
		Calculates the distance between two points using the distance formula.
		format is the output format you would like to use:
		dec: returns decimal (string) (ex. "2.82842712475")
		rad: returns a radical (string) (ex. "√8")
		simpRad: returns a simplified radical (string) (ex. "2√2")
	*/

	a := x2 - x1
	b := y2 - y1

	a = a * a
	b = b * b

	dist := math.Sqrt(a + b)

	if format == decimal {
		return strconv.FormatFloat(dist, 'f', -1, 64)
	}

	sqrtDist := "√" + strconv.FormatFloat(math.Round(dist*dist), 'f', -1, 64)

	if format == radical {
		return sqrtDist
	}

	root := math.Round(dist * dist)

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

	if format == simplifiedRadical {
		return simpleRoot
	}

	return "Error: format is not dec, rad, or simpRad"
}
