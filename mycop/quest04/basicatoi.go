package piscine

func BasicAtoi(s string) int {
	r := []rune(s)
	numb := 0
	lenght := 0
	for i := range r {
		lenght = i
	}
	on := 1
	for j := 0; j < lenght; j++ {
		on *= 10
	}
	for i := range r {
		nol := 0
		for k := '0'; k < r[i]; k++ {
			nol++
		}
		numb += nol * on
		on /= 10
	}
	return numb
}
