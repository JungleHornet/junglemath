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
