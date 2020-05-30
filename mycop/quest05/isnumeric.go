package piscine

func IsNumeric(str string) bool {
	num := true
	for i := range str {
		if str[i] > 47 && str[i] < 58 {
			num = true
		} else {
			return false
		}
	}
	return num
}
