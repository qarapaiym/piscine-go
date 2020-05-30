package main

import "github.com/01-edu/z01"

import "os"

func main() {
	word := os.Args
	lenght := 0
	for range word {
		lenght++
	}
	for i := lenght - 1; i > 0; i-- {
		if i == 0 {
			continue
		}
		runes := []rune(word[i])
		for j := range runes {
			z01.PrintRune(runes[j])
		}
		z01.PrintRune('\n')
	}
}
