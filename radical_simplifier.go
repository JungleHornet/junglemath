package junglemath

import (
	"math"
	"strconv"
)

func SimplifyRadical(root float64) string {

	rootCoefficient := int64(1)
	simpleRootInt := root

	for i := float64(2); i <= math.Round(math.Sqrt(root)); i++ {
		if (simpleRootInt / i) == math.Trunc(simpleRootInt/i) {
			if !(i == simpleRootInt) && !(i == 1) {
				if math.Sqrt(simpleRootInt/i) == math.Trunc(math.Sqrt(simpleRootInt/i)) {
					simpleRootInt = i
					rootCoefficient = rootCoefficient * int64(math.Sqrt(root/i))
				}
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

func CreateRoot(inpt float64) (string, bool) {

	root := inpt * inpt

	if root != math.Trunc(root) || root == 0 {
		return "", false
	}

	sqrtRoot := "√" + strconv.FormatFloat(root, 'f', -1, 64)

	return sqrtRoot, true
}
