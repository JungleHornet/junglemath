package junglemath

type Point struct {
	X, Y, Z float64
}

type Line struct {
	P1, P2 Point
}

type Triangle struct {
	A, B, C Point
}

type Angle struct {
	A, B, C Point
}

type Radical struct {
	Coef int64
	Rad  float64
}
