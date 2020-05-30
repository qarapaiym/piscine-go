package piscine

func IsAlpha(str string) bool {
	alph := true
	for i := range str {
		if (str[i] > 96 && str[i] < 123) || (str[i] > 64 && str[i] < 91) || (str[i] > 47 && str[i] < 58) || str[i] == 32 {
			alph = true
		} else {
			return false
		}
	}
	return alph
}
