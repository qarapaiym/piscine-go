
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args[1:]
	str := ar[0] + ar[1]
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
