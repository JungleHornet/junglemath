package junglemath

type Point struct {
	X, Y, Z float64
}

func (p *Point) Fix() {
	if p.X == 0 {
		p.X = 1
	}
	if p.Y == 0 {
		p.Y = 1
	}
	if p.Z == 0 {
		p.Z = 1
	}
}

type Line struct {
	P1, P2 Point
}

type Triangle struct {
	A, B, C Point
}

type Radical struct {
	Coef int64
	Rad  float64
}
