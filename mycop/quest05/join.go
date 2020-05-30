package piscine

func Join(strs []string, sep string) string {
	soz := ""
	first := false
	for _, c := range strs {
		if first {
			soz += sep
		}
		first = true
		soz += c
	}
	return soz
}
