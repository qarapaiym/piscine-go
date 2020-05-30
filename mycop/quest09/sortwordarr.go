package piscine

func SortWordArr(array []string) {
	arrayLen := -1
	for range array {
		arrayLen++
	}
	for i := 0; i < arrayLen; i++ {
		ind := i
		for j := i; j < arrayLen+1; j++ {
			if array[j] < array[ind] {
				ind = j
			}
		}
		array[i], array[ind] = array[ind], array[i]
	}
}
