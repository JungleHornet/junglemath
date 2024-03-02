package junglemath

import (
	"math"
)

func (t *Triangle) Orthocenter() Point {
	P1 := t.Centroid()
	P2 := t.Circumcenter()

	x := (3 * P1.X) - (2 * P2.X)
	y := (3 * P1.Y) - (2 * P2.Y)

	return Point{X: x, Y: y}
}

func (t *Triangle) Centroid() Point {
	x := (t.A.X + t.B.X + t.C.X) / 3
	y := (t.A.Y + t.B.Y + t.C.Y) / 3
	return Point{X: x, Y: y}
}

func (t *Triangle) Incenter() Point {
	A := Line{t.B, t.C}
	a := A.Length()
	B := Line{t.A, t.C}
	b := B.Length()
	C := Line{t.A, t.B}
	c := C.Length()

	x := (a*t.A.X + b*t.B.X + c*t.C.X) / (a + b + c)
	y := (a*t.A.Y + b*t.B.Y + c*t.C.Y) / (a + b + c)

	return Point{X: x, Y: y}
}

func (t *Triangle) Circumcenter() Point {
	x1 := t.A.X
	y1 := t.A.Y
	x2 := t.B.X
	y2 := t.B.Y
	x3 := t.C.X
	y3 := t.C.Y
	T := math.Pow(x1, 2) + math.Pow(y1, 2) - math.Pow(x2, 2) - math.Pow(y2, 2)
	u := math.Pow(x1, 2) + math.Pow(y1, 2) - math.Pow(x3, 2) - math.Pow(y3, 2)
	J := (x1-x2)*(y1-y3) - (x1-x3)*(y1-y2)

	x := ((-1*(y1-y2))*u + (y1-y3)*T) / (2 * J)
	y := ((x1-x2)*u - (x1-x3)*T) / (2 * J)

	return Point{X: x, Y: y}
}

func (a *Angle) Measure() float64 {
	A := Line{a.B, a.C}
	B := Line{a.A, a.C}
	C := Line{a.A, a.B}

	cc := C.Length()
	bb := B.Length()
	aa := A.Length()

	return math.Acos((math.Pow(bb, 2) + math.Pow(cc, 2) - math.Pow(aa, 2)) / (2 * bb * cc))
}
