package piscine

func IsPrintable(str string) bool {
	print := true
	for i := range str {
		if str[i] > 31 && str[i] < 128 {
			print = true
		} else {
			return false
		}
	}
	return print
}
