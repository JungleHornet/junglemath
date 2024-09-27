package junglemath

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/junglehornet/goscan"
)

func OpenCalculator() {
	/*
		Opens a calculator where you can solve equations with order of operations like 8 * (2/3 + 4).
		Enter "q" to exit.

		To Add: +
			Ex. 2+4 = 6
		To Subtract: -
			Ex. 2-4 = -2
		To Multiply: *
			Ex. 2*4 = 8
		To Divide: /
			Ex. 2/4 = 0.5
		To # Root: #r<number>
			Ex. 2r4 = 2

		This calculator environment is very basic and more of a template/example, so you are encouraged to make your own with some more polish, features, etc.
	*/
	s := goscan.NewScanner()
	inpt := s.ReadLine()
	first := true
	var ans string

	for inpt != "q" {
		if inpt != "" {
			if !strings.Contains(Digits, string([]rune(inpt)[0])) {
				if !first {
					inpt = ans + inpt
				}
			}
		}
		ans = strconv.FormatFloat(Solve(inpt), 'f', -1, 64)
		fmt.Println("= " + ans)
		first = false
		inpt = s.ReadLine()
	}
}

func PrepEquation(equation string) string {
	equation = "(" + equation + ")"
	equation = SolveOperator(equation, 0)
	re := regexp.MustCompile(`\((-?\d*\.?\d*)`)
	res := re.FindString(equation)
	newRes := strings.Replace(res, "-", "neg", -1)
	equation = strings.Replace(equation, res, newRes, -1)
	equation = strings.Replace(equation, "-", "min", -1)
	equation = strings.Replace(equation, "neg", "-", -1)
	equation = strings.TrimPrefix(equation, "(")
	equation = strings.TrimSuffix(equation, ")")
	return equation
}

func GetParentheses(inpt string) (string, bool) {
	if !strings.Contains(inpt, "(") {
		return "", false
	}
	var s scanner.Scanner
	s.Init(strings.NewReader(inpt))
	var inParantheses string
	var paranMode int
	paranMode = 0

	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		text := s.TokenText()
		if text == ")" {
			paranMode -= 1
			if paranMode == 0 {
				break
			}
		}
		if paranMode > 0 {
			inParantheses += text
		}

		if text == "(" {
			paranMode += 1
		}

	}
	return inParantheses, true
}

func SeparateVars(equation string, varName string) (string, float64) {
	equation = strings.ReplaceAll(equation, " ", "")
	equation = PrepEquation(equation)
	_, isParentheses := GetParentheses(equation)
	var bottom bool
	if isParentheses {
		bottom = false
	} else {
		bottom = true
	}
	var varAmount float64
	for !bottom {
	inParentheses, isParentheses := GetParentheses(equation)
	if !isParentheses {
		bottom = true
		break
	}
	innerEquation, innerVarAmount := SeparateVars(inParentheses, varName)
	index := strings.Index(equation, "("+inParentheses+")")
	innerSolution := 1.0
	coefficient := 1.0
	if index != -1 && index != 0 {
		coefString := ""
	for ;strings.Contains(Digits, string(equation[index-1])); index-- {
		coefString = string(equation[index-1]) + coefString
	}
	coefficient, _ = strconv.ParseFloat(coefString, 64)
	varAmount *= coefficient
	varAmount += innerVarAmount*coefficient
	innerSolution = Solve(innerEquation)
}
equation = strings.Replace(equation, strconv.FormatFloat(coefficient, 'f', -1, 64)+"("+inParentheses+")", strconv.FormatFloat(innerSolution*coefficient, 'f', -1, 64), -1)
}
re := regexp.MustCompile(`((min|\+?)\d*\.?\d*)` + varName)
for loop := 0; loop < 102; loop++ {
	if !strings.Contains(equation, varName) {
		break
	}
	match := re.FindStringSubmatch(equation)
	if len(match) == 3 && match[1] != match[2] {
	varCoef, err := strconv.ParseFloat(strings.Replace(strings.Replace(match[1], "min", "-", -1), "+", "", -1), 64)
	if err != nil {
		log.Fatal("Error: Invalid variablee " + varName + " passed to SeparateVars")
	}
	varAmount += varCoef
	equation = strings.Replace(equation, match[0], "", -1)
} else if match[1] == match[2] {
	varCoef := 1.0
	if match[1] == "+" {
		varCoef = 1.0
	} else if match[1] == "min" {
		varCoef = -1.0
	}
	fmt.Println(varAmount)
	varAmount += varCoef
	fmt.Println(varAmount)
	equation = strings.Replace(equation, match[0], "", -1)
} else {
	log.Fatal("Error: Invalid variable " + varName + " passed to SeparateVars")
}
}
equation = strings.Trim(equation, Operators)
return equation, varAmount
}

func SolveOperator(equation string, operator int) string {
	var opSymbol string
	var opRegex string
	switch operator {
	case 0:
		opRegex = `?r`
		opSymbol = "r"
	case 1:
		opRegex = `\^`
		opSymbol = "^"
	case 2:
		opRegex = `\*`
		opSymbol = "*"
	case 3:
		opRegex = `/`
		opSymbol = "/"
	case 4:
		opRegex = `\+`
		opSymbol = "+"
	case 5:
		opRegex = `min`
		opSymbol = "min"
	default:
		log.Fatal("Error: Invalid operator " + strconv.Itoa(operator) + " passed to SolveOperator")
	}
	Regex := regexp.MustCompile(`(-?\d*\.?\d*)` + opRegex + `(-?\d*\.?\d*)`)
	for {
		if !strings.Contains(equation, opSymbol) {
			break
		}
		operation := Regex.FindStringSubmatch(equation)
		num1, _ := strconv.ParseFloat(operation[1], 64)
		num2, _ := strconv.ParseFloat(operation[2], 64)
		var result string
		switch operator {
		case 0:
			result = strconv.FormatFloat(Root(num1, num2), 'f', -1, 64)
		case 1:
			result = strconv.FormatFloat(math.Pow(num1, num2), 'f', -1, 64)
		case 2:
			result = strconv.FormatFloat(num1*num2, 'f', -1, 64)
		case 3:
			result = strconv.FormatFloat(num1/num2, 'f', -1, 64)
		case 4:
			result = strconv.FormatFloat(num1+num2, 'f', -1, 64)
		case 5:
			result = strconv.FormatFloat(num1-num2, 'f', -1, 64)
		}
		equation = strings.Replace(equation, operation[0], result, -1)
	}
	return equation
}

func Solve(equation string) float64 {
	/*
		Solves an equation with order of operations like 8 * (2/3 + 4).
  
  		To Add: +
			Ex. 2+4 = 6
		To Subtract: -
			Ex. 2-4 = -2
		To Multiply: *
			Ex. 2*4 = 8
		To Divide: /
			Ex. 2/4 = 0.5
		To # Root: #r<number>
			Ex. 2r4 = 2
	*/
	equation = strings.ReplaceAll(equation, " ", "")
	equation = PrepEquation(equation)
	_, isParentheses := GetParentheses(equation)
	var solved bool
	if isParentheses {
		solved = false
	} else {
		solved = true
	}
	for !solved {
		inParentheses, _ := GetParentheses(equation)
		ans := Solve(inParentheses)
		equation = strings.Replace(equation, "("+inParentheses+")", strconv.FormatFloat(ans, 'f', -1, 64), -1)
		if !strings.Contains(equation, "(") {
			solved = true
		}
	}
	equation = SolveOperator(equation, 1)
	equation = SolveOperator(equation, 2)
	equation = SolveOperator(equation, 3)
	equation = SolveOperator(equation, 4)
	equation = SolveOperator(equation, 5)
	solved = false
	regex4 := regexp.MustCompile(`(-?\d*\.?\d*)min(-?\d*\.?\d*)`)
	for {
		if !strings.Contains(equation, "min") {
			break
		}
		minus := regex4.FindStringSubmatch(equation)
		num1, _ := strconv.ParseFloat(minus[1], 64)
		num2, _ := strconv.ParseFloat(minus[2], 64)
		result := strconv.FormatFloat(num1-num2, 'f', -1, 64)
		equation = strings.Replace(equation, minus[0], result, -1)
	}
	ans, err := strconv.ParseFloat(equation, 64)
	if err != nil {
		return 0
	}
	return ans
}

func SolveVar(equation string) float64 {
	/*
		Solves for x or another variable when given an equation such as 8(x-2) = 16 + x.
		Returns 0.0 if more than one = sign is present in the equation.

  		Same operators as Solve() are used.
	*/

	equations := strings.Split(equation, "=")
	if len(equations) > 2 {
		return 0.0
	}

	

	return 1
}