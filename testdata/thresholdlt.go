package main

import "fmt"

type Thresholdlt int

const (
	rlt1 Thresholdlt = 2
	rlt2 Thresholdlt = 4
	rlt3 Thresholdlt = 6
	rlt4 Thresholdlt = 8
	rlt5 Thresholdlt = 10
	rlt6 Thresholdlt = 12
	rlt7 Thresholdlt = 14
	rlt8 Thresholdlt = 16
	rlt9 Thresholdlt = 18
)

func main() {
	ck(1, "Thresholdlt(1)")
	ck(rlt1, "rlt1")
	ck(3, "Thresholdlt(3)")
	ck(rlt2, "rlt2")
	ck(5, "Thresholdlt(5)")
	ck(rlt3, "rlt3")
	ck(7, "Thresholdlt(7)")
	ck(rlt4, "rlt4")
	ck(9, "Thresholdlt(9)")
	ck(rlt5, "rlt5")
	ck(11, "Thresholdlt(11)")
	ck(rlt6, "rlt6")
	ck(13, "Thresholdlt(13)")
	ck(rlt7, "rlt7")
	ck(15, "Thresholdlt(15)")
	ck(rlt8, "rlt8")
	ck(17, "Thresholdlt(17)")
	ck(rlt9, "rlt9")
}

func ck(thresholdlt Thresholdlt, str string) {
	if fmt.Sprint(thresholdlt) != str {
		panic("thresholdlt.go: " + str)
	}
}
