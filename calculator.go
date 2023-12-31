package junglemath

import (
	"fmt"
	"github.com/junglehornet/goScan"
	"math"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
)

func OpenCalculator() {
	/*
		Opens a calculator where you can solve equations with order of operations like 8 * (2*3 + 4).
		Enter "q" to exit.
	*/
	s := goScan.NewScanner()
	var inpt string

	for inpt != "q" {
		inpt = s.ReadLine()
		fmt.Println(strconv.FormatFloat(Solve(inpt), 'f', -1, 64))
	}
	return
}

func Solve(equation string) float64 {
	/*
		Solves an equation with order of operations like 8 * (2*3 + 4).
	*/
	equation = strings.ReplaceAll(equation, " ", "")
	equation = PrepEquation(equation)
	inParentheses, isParentheses := GetParentheses(equation)
	var solved bool
	if isParentheses == false {
		solved = true
	} else {
		solved = false
	}
	for !solved {
		inParentheses, _ = GetParentheses(equation)
		ans := Solve(inParentheses)
		equation = strings.Replace(equation, "("+inParentheses+")", strconv.FormatFloat(ans, 'f', -1, 64), -1)
		if !strings.Contains(equation, "(") {
			solved = true
		}
	}
	solved = false
	expRegex := regexp.MustCompile("(-?\\d*\\.?\\d*)\\^(-?\\d*\\.?\\d*)")
	for !solved {
		if !strings.Contains(equation, "^") {
			solved = true
			break
		}
		exp := expRegex.FindStringSubmatch(equation)
		num1, _ := strconv.ParseFloat(exp[1], 64)
		num2, _ := strconv.ParseFloat(exp[2], 64)
		result := strconv.FormatFloat(math.Pow(num1, num2), 'f', -1, 64)
		equation = strings.Replace(equation, exp[0], result, -1)
	}
	solved = false
	regex1 := regexp.MustCompile("(-?\\d*\\.?\\d*)\\*(-?\\d*\\.?\\d*)")
	for !solved {
		if !strings.Contains(equation, "*") {
			solved = true
			break
		}
		mult := regex1.FindStringSubmatch(equation)
		num1, _ := strconv.ParseFloat(mult[1], 64)
		num2, _ := strconv.ParseFloat(mult[2], 64)
		result := strconv.FormatFloat(num1*num2, 'f', -1, 64)
		equation = strings.Replace(equation, mult[0], result, -1)
	}
	solved = false
	regex2 := regexp.MustCompile("(-?\\d*\\.?\\d*)/(-?\\d*\\.?\\d*)")
	for !solved {
		if !strings.Contains(equation, "/") {
			solved = true
			break
		}
		div := regex2.FindStringSubmatch(equation)
		num1, _ := strconv.ParseFloat(div[1], 64)
		num2, _ := strconv.ParseFloat(div[2], 64)
		result := strconv.FormatFloat(num1/num2, 'f', -1, 64)
		equation = strings.Replace(equation, div[0], result, -1)
	}
	solved = false
	regex3 := regexp.MustCompile("(-?\\d*\\.?\\d*)\\+(-?\\d*\\.?\\d*)")
	for !solved {
		if !strings.Contains(equation, "+") {
			solved = true
			break
		}
		plus := regex3.FindStringSubmatch(equation)
		num1, _ := strconv.ParseFloat(plus[1], 64)
		num2, _ := strconv.ParseFloat(plus[2], 64)
		result := strconv.FormatFloat(num1+num2, 'f', -1, 64)
		equation = strings.Replace(equation, plus[0], result, -1)
	}
	solved = false
	regex4 := regexp.MustCompile("(-?\\d*\\.?\\d*)min(-?\\d*\\.?\\d*)")
	for !solved {
		if !strings.Contains(equation, "min") {
			solved = true
			break
		}
		minus := regex4.FindStringSubmatch(equation)
		num1, _ := strconv.ParseFloat(minus[1], 64)
		num2, _ := strconv.ParseFloat(minus[2], 64)
		result := strconv.FormatFloat(num1-num2, 'f', -1, 64)
		equation = strings.Replace(equation, minus[0], result, -1)
	}
	ans, _ := strconv.ParseFloat(equation, 64)
	return ans
}

func PrepEquation(equation string) string {
	equation = "(" + equation + ")"
	re := regexp.MustCompile("\\((-?\\d*\\.?\\d*)")
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
