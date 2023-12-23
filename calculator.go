package junglemath

import (
	"fmt"
	"github.com/junglehornet/goScan"
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
		println(strconv.FormatFloat(Solve(inpt), 'f', -1, 64))
	}
	return
}

func Solve(equation string) float64 {
	/*
		Solves an equation with order of operations like 8 * (2*3 + 4).
	*/
	equation = strings.ReplaceAll(equation, " ", "")
	inParentheses := GetParentheses(equation)
	var solved bool
	if inParentheses == "" {
		solved = true
	} else {
		solved = false
	}
	for !solved {
		inParentheses = GetParentheses(equation)
		inParentheses = strings.TrimPrefix(inParentheses, "(")
		inParentheses = strings.TrimSuffix(inParentheses, ")")
		ans := Solve(inParentheses)
		fmt.Println(ans)
		equation = strings.Replace(equation, "("+inParentheses+")", strconv.FormatFloat(ans, 'f', -1, 64), -1)
		if !strings.Contains(equation, "(") {
			solved = true
		}
	}
	solved = false
	for !solved {
		regex1 := regexp.MustCompile("(-?\\d*\\.?\\d*)\\*(-?\\d*\\.?\\d*)")
		fmt.Println(equation)
		mult := regex1.FindStringSubmatch(equation)
		fmt.Println(mult)
		mult1, _ := strconv.ParseFloat(mult[1], 64)
		mult2, _ := strconv.ParseFloat(mult[2], 64)
		multRes := strconv.FormatFloat(mult1*mult2, 'f', -1, 64)
		equation = strings.Replace(equation, mult[0], multRes, -1)
		if !strings.Contains(equation, "*") {
			solved = true
		}
	}
	ans, _ := strconv.ParseFloat(equation, 64)
	return ans
}

func GetParentheses(inpt string) string {
	var s scanner.Scanner
	s.Init(strings.NewReader(inpt))
	var inParantheses string
	var paranMode int
	paranMode = 0

	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		if token == ')' {
			paranMode -= 1
			if paranMode == 0 {
				break
			}
		}
		if paranMode > 0 {
			inParantheses += s.TokenText()
		}

		if token == '(' {
			paranMode += 1
		}

	}
	return inParantheses
}
