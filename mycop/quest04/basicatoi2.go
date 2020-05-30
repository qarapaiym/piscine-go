package piscine

func BasicAtoi2(s string) int {
	r := []rune(s)
	num := 0
	lenght := 0
	for i := range r {
		lenght = i
	}
	on := 1
	for j := 0; j < lenght; j++ {
		on *= 10
	}
	for i := range r {
		if r[i] < '0' || r[i] > '9' {
			return 0
		}
		tmp := 0
		for k := '0'; k < r[i]; k++ {
			tmp++
		}
		num += tmp * on
		on /= 10
	}
	return num
}
