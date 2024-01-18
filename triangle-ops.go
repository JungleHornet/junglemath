package junglemath

import "math"

func GetOrthocenter(x1, y1, x2, y2, x3, y3 float64) (float64, float64) {
	X1, Y1 := GetCentroid(x1, y1, x2, y2, x3, y3)
	X2, Y2 := GetCircumcenter(x1, y1, x2, y2, x3, y3)
	
	x := (3 * X1) - (2 * X2)
	y := (3 * Y1) - (2 * Y2)

	return x, y
}

func GetCentroid(x1, y1, x2, y2, x3, y3 float64) (float64, float64) {
	x := (x1 + x2 + x3) / 3
	y := (y1 + y2 + y3) / 3
	return x, y
}

func GetCircumcenter(x1, y1, x2, y2, x3, y3 float64) (float64, float64) {
	t := math.Pow(x1, 2) + math.Pow(y1, 2) - math.Pow(x2, 2) - math.Pow(y2, 2)
	u := math.Pow(x1, 2) + math.Pow(y1, 2) - math.Pow(x3, 2) - math.Pow(y3, 2)
	J := (x1 - x2) * (y1 - y3) - (x1 - x3) * (y1 - y2)

	x := ((-1 * (y1 - y2)) * u + (y1 - y3) * t) / (2 * J)
	y := ((x1 - x2) * u - (x1 - x3) * t) / (2 * J)

	return x, y
}

func GetIncenter(x1, y1, x2, y2, x3, y3 float64) (float64, float64) {
	a := CalcDistance(x1, y1, x2, y2)
	b := CalcDistance(x2, y2, x3, y3)
	c := CalcDistance(x3, y3, x1, y1)

	x := (a * x1 + b * x2 + c * x3) / (a + b + c)
	y := (a * y1 + b * y2 + c * y3) / (a + b + c)

	return x, y
}