package piscine

func BasicJoin(strs []string) string {
	soz := ""
	for i := range strs {
		soz = soz + strs[i]
	}
	return soz
}
