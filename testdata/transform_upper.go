package main

import "fmt"

type UpperCaseValue int

const (
	UpperCaseValueOne UpperCaseValue = iota
	UpperCaseValueTwo
	UpperCaseValueThree
)

func main() {
	ck(UpperCaseValueOne, "UPPERCASEVALUEONE")
	ck(UpperCaseValueTwo, "UPPERCASEVALUETWO")
	ck(UpperCaseValueThree, "UPPERCASEVALUETHREE")
	ck(-127, "")
	ck(127, "")
}

func ck(value UpperCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_upper.go: " + str)
	}
}
