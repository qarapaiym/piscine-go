package piscine

func LastRune(s string) rune {
	sent := []rune(s)
	l := 0
	for index := range sent {
		l = index
	}
	return sent[(l)]
}
