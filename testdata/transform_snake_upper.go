package main

import "fmt"

type SnakeUpperCaseValue int

const (
	SnakeUpperCaseValueOne SnakeUpperCaseValue = iota
	SnakeUpperCaseValueTwo
	SnakeUpperCaseValueThree
)

func main() {
	ck(SnakeUpperCaseValueOne, "SNAKE_UPPER_CASE_VALUE_ONE")
	ck(SnakeUpperCaseValueTwo, "SNAKE_UPPER_CASE_VALUE_TWO")
	ck(SnakeUpperCaseValueThree, "SNAKE_UPPER_CASE_VALUE_THREE")
	ck(-127, "")
	ck(127, "")
}

func ck(value SnakeUpperCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_snake_upper.go: " + str)
	}
}
