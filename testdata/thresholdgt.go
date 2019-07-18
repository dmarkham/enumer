package main

import "fmt"

type Thresholdgt int

const (
	rgt1  Thresholdgt = 2
	rgt2  Thresholdgt = 4
	rgt3  Thresholdgt = 6
	rgt4  Thresholdgt = 8
	rgt5  Thresholdgt = 10
	rgt6  Thresholdgt = 12
	rgt7  Thresholdgt = 14
	rgt8  Thresholdgt = 16
	rgt9  Thresholdgt = 18
	rgt10 Thresholdgt = 20
	rgt11 Thresholdgt = 22
)

func main() {
	ck(1, "Thresholdgt(1)")
	ck(rgt1, "rgt1")
	ck(3, "Thresholdgt(3)")
	ck(rgt2, "rgt2")
	ck(5, "Thresholdgt(5)")
	ck(rgt3, "rgt3")
	ck(7, "Thresholdgt(7)")
	ck(rgt4, "rgt4")
	ck(9, "Thresholdgt(9)")
	ck(rgt5, "rgt5")
	ck(11, "Thresholdgt(11)")
	ck(rgt6, "rgt6")
	ck(13, "Thresholdgt(13)")
	ck(rgt7, "rgt7")
	ck(15, "Thresholdgt(15)")
	ck(rgt8, "rgt8")
	ck(17, "Thresholdgt(17)")
	ck(rgt9, "rgt9")
	ck(19, "Thresholdgt(19)")
	ck(rgt10, "rgt10")
	ck(21, "Thresholdgt(21)")
	ck(rgt11, "rgt11")
}

func ck(thresholdgt Thresholdgt, str string) {
	if fmt.Sprint(thresholdgt) != str {
		panic("thresholdgt.go: " + str)
	}
}