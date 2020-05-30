package piscine

func Any(f func(string) bool, arr []string) bool {
	b := false
	for i := range arr {
		if f(arr[i]) == true {
			b = true
		}
	}
	return b
}
