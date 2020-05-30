package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(table []string) {
	for i := range table {
		for j := range table[i] {
			z01.PrintRune(rune(table[i][j]))
		}
		z01.PrintRune(10)
	}
}
