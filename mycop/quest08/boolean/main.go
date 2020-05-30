package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	os := os.Args[1:]
	lengthOfArg := 0
	for range os {
		lengthOfArg++
	}
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

func printStr(str string) {
	arrayStr := []rune(str)
	for i := range arrayStr {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(s int) bool {
	if s%2 != 1 {
		return true
	} else {
		return false
	}
}
