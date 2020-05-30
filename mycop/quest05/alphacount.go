package piscine

func AlphaCount(str string) int {
	r := []rune(str)
	c := 0
	for i := range r {
		if (r[i] > 64 && r[i] < 91) || (r[i] > 96 && r[i] < 123) {
			c++
		}
	}
	return c
}
