package main

import "os"

import "github.com/01-edu/z01"

func main() {
	arg := os.Args
	argRune := []rune(arg[0])
	z01.PrintRune('.')
	z01.PrintRune('/')
	for i := range argRune {
		z01.PrintRune(argRune[i])
	}
	z01.PrintRune(10)
}
