package piscine

func Unmatch(arr []int) int {
	len := 0
	for range arr {
		len++
	}
	for i := range arr {
		count := 0
		for j := 0; j < len; j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		if count%2 == 1 {
			return arr[i]
		}
	}
	return -1
}
