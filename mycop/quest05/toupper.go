package piscine

func ToUpper(s string) string {
	str := []byte(s)
	for i, j := range str {
		if j > 96 && j < 123 {
			str[i] = str[i] - 32
		}
	}
	return string(str)
}
