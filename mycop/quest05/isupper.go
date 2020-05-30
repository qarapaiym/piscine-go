package piscine

func IsUpper(str string) bool {
	up := true
	for i := range str {
		if str[i] >= 'A' && str[i] <= 'Z' {
			up = true
		} else {
			return false
		}
	}
	return up
}
