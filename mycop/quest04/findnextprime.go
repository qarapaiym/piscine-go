package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	for nb > 0 {
		if IsPrime2(nb) {
			return nb
		}
		nb++
	}
	return nb
}

func IsPrime2(nb int) bool {

	if nb < 0 {
		for i := -2; i >= nb/2; i-- {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
	if nb == 0 || nb == -1 || nb == 1 {
		return false
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
