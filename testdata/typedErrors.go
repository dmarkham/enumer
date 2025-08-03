package main

import "fmt"

type TypedErrorsValue int

const (
	TypedErrorsValueOne TypedErrorsValue = iota
	TypedErrorsValueTwo
	TypedErrorsValueThree
)

func main() {
	checkMatch(TypedErrorsValueOne, "TypedErrorsValueOne")
	checkMatch(TypedErrorsValueTwo, "TypedErrorsValueTwo")
	checkMatch(TypedErrorsValueThree, "TypedErrorsValueThree")
	checkMatch(-127, "TypedErrorsValue(-127)")
	checkMatch(127, "TypedErrorsValue(127)")
}

func checkMatch(value TypedErrorsValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_upper.go: " + str)
	}
}
