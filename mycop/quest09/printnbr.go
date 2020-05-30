package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	last := false
	if n == -9223372036854775808 {
		n = -922337203685477580
		last = true
	}
	if n < 0 {
		n = n * -1
		z01.PrintRune('-')
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	var myNum [36]int
	first := n
	sec := 0
	i := 0
	for first > 0 {
		sec = first % 10
		first /= 10
		myNum[i] = sec
		i++
	}
	for j := i - 1; j >= 0; j-- {
		l := 0
		varik := '0'
		for l < myNum[j] {
			varik++
			l++
		}
		z01.PrintRune(varik)
	}
	if last == true {
		z01.PrintRune('8')
	}
}
