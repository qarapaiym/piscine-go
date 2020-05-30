package piscine

func IsPrime(nb int) bool {
	res := true
	if nb == 0 || nb == 1 || nb == -1 {
		return false
	}
	if nb > 0 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				res = false
			}
		}
	} else {
		for i := -2; i > nb; i-- {
			if nb%i == 0 {
				res = false
			}
		}
	}
	return res
}
