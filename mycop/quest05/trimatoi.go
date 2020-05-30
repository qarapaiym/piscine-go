package piscine

func TrimAtoi(s string) int {
	array := []int{}
	for i := range s {
		if s[i] > 48 && s[i] < 58 || s[i] == 45 {
			array = append(array, int(s[i]))
		}
	}
	arrTwo := []int{}
	for i := range array {
		if array[i] != 45 || array[0] == 45 {
			arrTwo = append(arrTwo, array[i])
		}
	}
	length := 0
	for range arrTwo {
		length++
	}
	soz := ""
	for i := 0; i < length; i++ {
		soz = soz + string(arrTwo[i])
	}
	r := []rune(soz)
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
	for range r {
		if r[0] == '-' {
			numb = numb * -1
		}
	}
	return numb
}