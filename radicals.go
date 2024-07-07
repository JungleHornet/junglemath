package junglemath

import (
	"math"
	"strconv"
)

func (r *Radical) Simplify() string {

	rootCoefficient := r.Coef
	simpleRootInt := r.Rad
	root := r.Rad

	for i := float64(2); i <= (math.Ceil(math.Sqrt(root))); i++ {
		if (simpleRootInt / i) == math.Trunc(simpleRootInt/i) {
			if !(i == simpleRootInt) {
				if math.Sqrt(simpleRootInt/i) == math.Trunc(math.Sqrt(simpleRootInt/i)) {
					simpleRootInt = i
					rootCoefficient = rootCoefficient * int64(math.Sqrt(root/i))
				} else if math.Sqrt(i) == math.Trunc(math.Sqrt(i)) {
					simpleRootInt = simpleRootInt / i
					rootCoefficient = rootCoefficient * int64(math.Sqrt(i))
				}
				i = 2
			}
		}
	}

	var simpleRoot string
	if rootCoefficient == 1 {
		simpleRoot = "√" + strconv.FormatFloat(root, 'f', -1, 64)
	} else {
		simpleRoot = strconv.FormatInt(rootCoefficient, 10) + "√" +
			strconv.FormatFloat(simpleRootInt, 'f', -1, 64)
	}

	return simpleRoot

}

func CreateRoot(inpt float64) (Radical, bool) {

	root := inpt * inpt

	if root != math.Trunc(root) || root == 0 {
		return Radical{}, false
	}

	sqrtRoot := Radical{1, root}

	return sqrtRoot, true
}

func Root(x, y float64) float64 {
	/*
		Returns x√y
	*/
	exp := 1 / x
	ans := math.Pow(y, exp)
	return ans
}
