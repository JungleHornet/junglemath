package junglemath

import (
	"log"
	"reflect"
)

func IsValidPointMap(m map[string]any) bool {
	delete(m, "type")
	if reflect.TypeOf(m["X"]).Kind() != reflect.Float64 {
		return false
	}
	if reflect.TypeOf(m["Y"]).Kind() != reflect.Float64 {
		return false
	}
	if reflect.TypeOf(m["Z"]).Kind() != reflect.Float64 {
		return false
	}
	return true
}

func ToPoint(m map[string]any, name string) Point {
	if IsValidPointMap(m) {
		X := m["X"].(float64)
		Y := m["Y"].(float64)
		Z := m["Z"].(float64)
		point := Point{X: X, Y: Y, Z: Z}
		return point
	}
	log.Fatal("\033[1;31mError: Invalid stored variable: " + name + ". Please use jcalc -set to reset the variable's values.\033[0m")
	return Point{}
}

func IsValidLineMap(m map[string]any) bool {
	P1, success := m["P1"].(map[string]any)
	if !success {
		return false
	}
	if !IsValidPointMap(P1) {
		return false
	}
	P2, success := m["P2"].(map[string]any)
	if !success {
		return false
	}
	if !IsValidPointMap(P2) {
		return false
	}
	return true
}

func ToLine(m map[string]any, name string) Line {
	if IsValidLineMap(m) {
		P1 := ToPoint(m["P1"].(map[string]any), name)
		P2 := ToPoint(m["P2"].(map[string]any), name)
		line := Line{P1: P1, P2: P2}
		return line
	}
	log.Fatal("\033[1;31mError: Invalid stored variable: " + name + ". Please use jcalc -set to reset the variable's values.\033[0m")
	return Line{}
}

func IsValidAngleMap(m map[string]any) bool {
	pointA, success := m["A"].(map[string]any)
	if !success {
		return false
	}
	if !IsValidPointMap(pointA) {
		return false
	}
	pointB, success := m["B"].(map[string]any)
	if !success {
		return false
	}
	if !IsValidPointMap(pointB) {
		return false
	}
	pointC, success := m["C"].(map[string]any)
	if !success {
		return false
	}
	if !IsValidPointMap(pointC) {
		return false
	}
	return true
}

func ToAngle(m map[string]any, name string) Angle {
	if IsValidAngleMap(m) {
		pointA := ToPoint(m["A"].(map[string]any), name)
		pointB := ToPoint(m["B"].(map[string]any), name)
		pointC := ToPoint(m["C"].(map[string]any), name)
		angle := Angle{A: pointA, B: pointB, C: pointC}
		return angle
	}
	log.Fatal("\033[1;31mError: Invalid stored variable: " + name + ". Please use jcalc -set to reset the variable's values.\033[0m")
	return Angle{}
}

func IsValidTriangleMap(m map[string]any) bool {
	pointA, success := m["A"].(map[string]any)
	if !success {
		return false
	}
	if !IsValidPointMap(pointA) {
		return false
	}
	pointB, success := m["B"].(map[string]any)
	if !success {
		return false
	}
	if !IsValidPointMap(pointB) {
		return false
	}
	pointC, success := m["C"].(map[string]any)
	if !success {
		return false
	}
	if !IsValidPointMap(pointC) {
		return false
	}
	return true
}

func ToTriangle(m map[string]any, name string) Triangle {
	if IsValidTriangleMap(m) {
		pointA := ToPoint(m["A"].(map[string]any), name)
		pointB := ToPoint(m["B"].(map[string]any), name)
		pointC := ToPoint(m["C"].(map[string]any), name)
		angle := Triangle{A: pointA, B: pointB, C: pointC}
		return angle
	}
	log.Fatal("\033[1;31mError: Invalid stored variable: " + name + ". Please use jcalc -set to reset the variable's values.\033[0m")
	return Triangle{}
}

func IsValidTriangle(t Triangle) bool {
	angleA := Angle{A: t.B, B: t.A, C: t.C}
	angleB := Angle{A: t.A, B: t.B, C: t.C}
	angleC := Angle{A: t.A, B: t.C, C: t.B}
	if angleA.Measure()+angleB.Measure()+angleC.Measure() != 180 {
		return false
	}
	aSide := Line{P1: t.B, P2: t.C}
	bSide := Line{P1: t.A, P2: t.C}
	cSide := Line{P1: t.A, P2: t.B}
	if aSide.Length()+bSide.Length() <= cSide.Length() {
		return false
	}
	if aSide.Length()+cSide.Length() <= bSide.Length() {
		return false
	}
	if bSide.Length()+cSide.Length() <= aSide.Length() {
		return false
	}
	return true
}
