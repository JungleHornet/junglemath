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

package main

import (
	"fmt"
	"github.com/junglehornet/junglemath"
)

func main() {
	// good test input for Solve(): 2r4-(15/(2+1)+1)*2^2
	rad := junglemath.Radical{Rad: 7224653312}
	fmt.Println(rad.Simplify())
}
