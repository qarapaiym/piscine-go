package piscine

func IsLower(str string) bool {
	low := true
	for i := range str {
		if str[i] > 96 && str[i] < 123 {
			low = true
		} else {
			return false
		}
	}
	return low
}
