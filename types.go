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
