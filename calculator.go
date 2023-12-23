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
	regex1 := regexp.MustCompile("\\(.*\\)")
	equation = strings.ReplaceAll(equation, " ", "")
	inParentheses := regex1.FindString(equation)
	var solved bool
	if inParentheses == "" {
		solved = true
	} else {
		solved = false
	}
	for !solved {
		inParentheses = regex1.FindString(equation)
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
		regex2 := regexp.MustCompile("(-?\\d*\\.?\\d*)\\*(-?\\d*\\.?\\d*)")
		fmt.Println(equation)
		mult := regex2.FindStringSubmatch(equation)
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

func GetParentheses(inpt string) {
	var s scanner.Scanner
	s.Init(strings.NewReader(inpt))

	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		text := s.TokenText()
		fmt.Println(token)
		fmt.Println(text)
	}
}
