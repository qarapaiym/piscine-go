package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x string
	y string
}

func main() {
	points := point{}

	points.x = "42"
	points.y = "21"
	firstStr := "x = "
	secStr := ", y = "
	full := firstStr + points.x + secStr + points.y
	for _, c := range full {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
