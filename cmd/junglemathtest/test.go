package main

import (
	"fmt"
	"strconv"

	"github.com/junglehornet/junglemath"
)

func main() {
	// good test input for Solve(): 8-(15/(2+1)+1)*2^2
	x, y := junglemath.GetCentroid(2, 2, 1, 1, 3, 1)
	xstr := strconv.FormatFloat(x, 'f', -1, 64)
	ystr := strconv.FormatFloat(y, 'f', -1, 64)
	fmt.Println("(" + xstr + "," + ystr + ")")
}
