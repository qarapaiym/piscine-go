
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args[1:]
	first := os.Args[1]
	sec := os.Args[2]
	str := ""
	if len(ar) == 2 {
		for i := range first {
			for j := range sec {
				if first[i] == sec[j] {
					str += string(first[i])
				}
			}
		}
	}
	newStr := ""
	bul := false
	for i := range str {
		bul = false
		for j := range newStr {
			if str[i] == newStr[j] {
				bul = true
			}
		}
		if !bul {
			newStr += string(str[i])
		}
	}
	for _, z := range newStr {
		z01.PrintRune(z)
	}
	z01.PrintRune('\n')
}
