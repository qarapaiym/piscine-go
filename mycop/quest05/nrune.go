package piscine

func NRune(s string, n int) rune {
	sent := []rune(s)
	for i, l := range sent {
		if n-1 == i {
			return l
		}
	}
	return 0
}
