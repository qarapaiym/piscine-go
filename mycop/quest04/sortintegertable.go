package piscine

func SortIntegerTable(table []int) {
	tableLen := len(table) - 1
	for i := 0; i < tableLen; i++ {
		ind := i
		for j := i; j < tableLen+1; j++ {
			if table[j] < table[ind] {
				ind = j
			}
		}
		table[i], table[ind] = table[ind], table[i]
	}
}
