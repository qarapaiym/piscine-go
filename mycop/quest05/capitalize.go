package piscine

func Capitalize(s string) string {
	r := []rune(s)
	soz := false
	for i := range r {
		if r[i] > 96 && r[i] < 123 && soz == false {
			r[i] = r[i] - 32
			soz = true
		} else if r[i] > 64 && r[i] < 91 && soz == false {
			soz = true
		} else if r[i] > 64 && r[i] < 91 && soz == true {
			r[i] = r[i] + 32
			soz = true
		} else if r[i] > 96 && r[i] < 123 && soz == true {
			soz = true
		} else if r[i] > 47 && r[i] < 58 {
			soz = true
		} else {
			soz = false
		}
	}
	return string(r)
}
