package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	len := 0
	for range tab {
		len++
	}
	left := false
	right := false
	for i := 0; i < len-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			left = true
			println("adas")
		} else if f(tab[i], tab[i+1]) < 0 {
			right = true
			println("665465")
		}
	}
	if left && right {
		return false
	} else {
		return true
	}
}
