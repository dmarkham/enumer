package main

import "fmt"

type LowerCaseValue int

const (
	LowerCaseValueOne LowerCaseValue = iota
	LowerCaseValueTwo
	LowerCaseValueThree
)

func main() {
	ck(LowerCaseValueOne, "lowercasevalueone")
	ck(LowerCaseValueTwo, "lowercasevaluetwo")
	ck(LowerCaseValueThree, "lowercasevaluethree")
	ck(-127, "")
	ck(127, "")
}

func ck(value LowerCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_lower.go: " + str)
	}
}
