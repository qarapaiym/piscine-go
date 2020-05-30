package piscine

func MakeRange(min, max int) []int {
	var null []int
	if min >= max {
		return null
	}
	size := max - min
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i + min
	}
	return arr
}
