package piscine

func Map(f func(int) bool, arr []int) []bool {
	len := 0
	for range arr {
		len++
	}
	array := make([]bool, len)

	for i, c := range arr {
		array[i] = f(c)
	}
	return array
}
