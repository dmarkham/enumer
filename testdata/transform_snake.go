package main

import "fmt"

type SnakeCaseValue int

const (
	SnakeCaseValueOne SnakeCaseValue = iota
	SnakeCaseValueTwo
	SnakeCaseValueThree
)

func main() {
	ck(SnakeCaseValueOne, "snake_case_value_one")
	ck(SnakeCaseValueTwo, "snake_case_value_two")
	ck(SnakeCaseValueThree, "snake_case_value_three")
	ck(-127, "")
	ck(127, "")
}

func ck(value SnakeCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_snake.go: " + str)
	}
}
