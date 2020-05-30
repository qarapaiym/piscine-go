package piscine

func Index(s string, toFind string) int {
	countOne := 0
	countTwo := 0
	for i := range s {
		countOne = i
	}
	for i := range toFind {
		countTwo = i
	}
	if toFind == "" {
		return 0
	}
	for i := 0; i < countOne-countTwo; i++ {
		if s[i:i+countTwo+1] == toFind {
			return i
		}
	}
	return -1
}
