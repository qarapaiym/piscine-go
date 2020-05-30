package piscine

func StrRev(s string) string {
	var str string
	for _, rng := range s {
		str = string(rng) + str
	}
	return str
}
