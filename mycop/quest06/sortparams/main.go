package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	sorting := os.Args
	var sortLen int
	for range sorting {
		sortLen++
	}
	for i := 0; i < sortLen; i++ {
		ind := i
		for j := i; j < sortLen; j++ {
			if sorting[j] < sorting[ind] {
				ind = j
			}
		}
		sorting[i], sorting[ind] = sorting[ind], sorting[i]
	}
	for i := 0; i < sortLen-1; i++ {
		runes := []rune(sorting[i])
		for j := range runes {
			z01.PrintRune(runes[j])
		}
		z01.PrintRune('\n')
	}
}
