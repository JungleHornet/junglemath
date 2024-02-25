package junglemath

import (
	"math"
)

func (l *Line) length() float64 {
	/*
		Calculates the distance between two points using the distance formula.
	*/

	x1 := l.P1.X
	y1 := l.P1.Y
	z1 := l.P1.Z
	x2 := l.P2.X
	y2 := l.P2.Y
	z2 := l.P2.Z

	a := x2 - x1
	b := y2 - y1
	c := z2 - z1

	a = a * a
	b = b * b
	c = c * c

	dist := math.Sqrt(a + b + c)

	return dist
}
