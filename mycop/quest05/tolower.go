package piscine

func ToLower(s string) string {
	str := []byte(s)
	for i, j := range str {
		if j > 64 && j < 91 {
			str[i] = str[i] + 32
		}
	}
	return string(str)
}
