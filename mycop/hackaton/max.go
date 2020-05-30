package piscine

func Max(arr []int) int {
	len := len(arr)

	for i := 0; i < len-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			i = -1
		}

	}
	return arr[len-1]
}
