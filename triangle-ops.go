package junglemath

import "math"

func GetOrthocenter(x1, y1, x2, y2, x3, y3 float64) (float64, float64) {
	X1, Y1 := GetCentroid(x1, y1, x2, y2, x3, y3)
	X2, Y2 := GetCircumcenter(x1, y1, x2, y2, x3, y3)
	
	x := (3 * X2) - (2 * X1) / (3 - 2)
	y := (3 * Y2) - (2 * Y1) / (3 - 2)

	return x, y
}

func GetCentroid(x1, y1, x2, y2, x3, y3 float64) (float64, float64) {
	x := (x1 + x2 + x3) / 3
	y := (y1 + y2 + y3) / 3
	return x, y
}

func GetCircumcenter(x1, y1, x2, y2, x3, y3 float64) (float64, float64) {
	a := CalcDistance(x1, y1, x2, y2)
	b := CalcDistance(x2, y2, x3, y3)
	c := CalcDistance(x3, y3, x1, y1)
	A := math.Acos((math.Pow(a, 2) + math.Pow(b, 2) - math.Pow(c, 2)) / (2 * b * c))
	B := math.Acos((math.Pow(a, 2) + math.Pow(b, 2) - math.Pow(c, 2)) / (2 * a * c))
	C := math.Acos((math.Pow(a, 2) + math.Pow(b, 2) - math.Pow(c, 2)) / (2 * a * b))

	x := (x1 * math.Sin(2 * A) + x2 * math.Sin(2 * B) + x3 * math.Sin(2 * C)) / (math.Sin(2 * A) + math.Sin(2 * B) + math.Sin(2 * C))
	y := (y1 * math.Sin(2 * A) + y2 * math.Sin(2 * B) + y3 * math.Sin(2 * C)) / (math.Sin(2 * A) + math.Sin(2 * B) + math.Sin(2 * C))
	return x, y
}