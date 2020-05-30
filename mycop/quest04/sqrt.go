package piscine

func Sqrt(nb int) int {
	res := 0
	for i := 0; i <= nb; i++ {
		if i*i == nb {
			res = i
		}
	}
	return res
}
