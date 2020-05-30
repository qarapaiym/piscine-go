package piscine

func ForEach(f func(int), arr []int) {
	for _, c := range arr {
		f(c)
	}
}
