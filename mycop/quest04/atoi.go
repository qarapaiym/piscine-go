package piscine

func Atoi(s string) int {
	r := []rune(s)
	num := 0
	lenght := 0
	for i := range r {
		lenght = i
	}
	if lenght == 0 {
		return 0
	}
	on := 1
	for j := 0; j < lenght; j++ {
		if r[j] == '+' || r[j] == '-' {
			continue
		}
		on *= 10
	}
	for i := range r {
		if (r[0] == '+' || r[0] == '-') && i == 0 {
			continue
		}
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
	if r[0] == '-' {
		num *= -1
	}
	return num
}
