package junglemath

import (
	"fmt"
	"math"
	"strconv"
)

func CalcDistance(x1 float64, y1 float64, x2 float64, y2 float64) string {

	/*
		Calculates the distance between two points using the distance formula.
		format is the output format you would like to use:
		dec: returns Decimal (string) (ex. "2.82842712475")
		rad: returns a Radical (string) (ex. "√8")
		simpRad: returns a simplified Radical (string) (ex. "2√2")
	*/

	a := x2 - x1
	b := y2 - y1

	a = a * a
	b = b * b

	dist := math.Sqrt(a + b)

	return strconv.FormatFloat(dist, 'f', -1, 64)

}

func CalcDistance3D(x1, y1, z1, x2, y2, z2 float64) string {
	twoDimDist, err := strconv.ParseFloat(CalcDistance(x1, y1, x2, y2), 64)
	if err != nil {
		fmt.Println(err)
	}

	twoDimDist = math.Round(twoDimDist * twoDimDist)
	c := z2 - z1

	c = c * c

	dist := math.Sqrt(twoDimDist + c)

	return strconv.FormatFloat(dist, 'f', -1, 64)
}
