package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	array := []int{}
	if n == 0 {
		z01.PrintRune('0')
	}
	for n != 0 {
		float := n % 10
		array = append(array, float)
		n = n / 10
	}
	length := 0
	for range array {
		length++
	}

	tableLen := length - 1

	for i := 0; i < tableLen; i++ {
		ind := i
		for j := i; j < tableLen+1; j++ {
			if array[j] < array[ind] {
				ind = j
			}
		}
		array[i], array[ind] = array[ind], array[i]
	}

	for _, i := range array {
		z01.PrintRune(rune(i + 48))
	}
}
