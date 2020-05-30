package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	count := 0
	for range array {
		count++
	}
	for i := 0; i < count-1; i++ {
		if f(array[i], array[i+1]) == 1 {
			array[i], array[i+1] = array[i+1], array[i]
			i = -1
		}
	}
}

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
