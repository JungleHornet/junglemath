package main

import (
	"fmt"

	"github.com/junglehornet/junglemath"
)

func main() {
	// good test input for Solve(): 2r4-(15/(2+1)+1)*2^2
	junglemath.OpenCalculator()
	fmt.Println("test")
	fmt.Println(junglemath.SeparateVars("2x + 2(4 - 2x)", "x"))
}
