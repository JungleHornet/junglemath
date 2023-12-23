package junglemath

import (
	"fmt"
	"github.com/junglehornet/goScan"
	"log"
	"regexp"
	"strconv"
	"strings"
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
	if inParentheses != "" {
		inParentheses = strings.TrimLeft(inParentheses, "(")
		inParentheses = strings.TrimRight(inParentheses, ")")
		ans := Solve(inParentheses)
		equation = strings.Replace(equation, "("+inParentheses+")", strconv.FormatFloat(ans, 'f', -1, 64), -1)
	} else {
		regex2 := regexp.MustCompile("(-?\\d*.?\\d*)\\*(-?\\d*.?\\d*)")
		for mult := regex2.FindStringSubmatch(equation); mult != nil; mult = regex2.FindStringSubmatch(equation) {
			mult1, _ := strconv.ParseFloat(mult[1], 64)
			mult2, _ := strconv.ParseFloat(mult[2], 64)
			multRes := strconv.FormatFloat(mult1*mult2, 'f', -1, 64)
			equation = strings.Replace(equation, mult[0], multRes, -1)
		}
	}
	ans, err := strconv.ParseFloat(equation, 64)
	fmt.Println(ans)
	if err != nil {
		fmt.Println("Error: Unable to parse final equation.")
		log.Fatal(err)
	}
	return ans
}
