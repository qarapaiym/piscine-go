package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	words := os.Args[1:]
	isUp := false
	for _, v := range words {
		if v == "--upper" {
			isUp = true
		}
	}
	for _, val := range words {
		if val == "--upper" {
			continue
		}
		num := 0
		for _, v := range val {
			num = num*10 + int(v-'0')
		}
		if num > 0 && num < 27 {
			if isUp == false {
				z01.PrintRune(rune(num + 96))
			} else {
				z01.PrintRune(rune(num + 64))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
