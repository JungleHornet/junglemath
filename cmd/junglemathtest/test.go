package main

import (
	"fmt"
	"strconv"

	"github.com/junglehornet/junglemath"
)

func main() {
	// good test input for Solve(): 8-(15/(2+1)+1)*2^2
	x, y := junglemath.GetCircumcenter(6, 0, 0, 0, 0, 8)
	xstr := strconv.FormatFloat(x, 'f', -1, 64)
	ystr := strconv.FormatFloat(y, 'f', -1, 64)
	fmt.Println("(" + xstr + "," + ystr + ")")
}
