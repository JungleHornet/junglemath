/*
	JungleMath - A Go library for advanced Math operations
    Copyright (C) 2024  JungleHornet

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package junglemath

import (
	"math"
)

func (l *Line) Length() float64 {
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
