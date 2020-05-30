package piscine

func Compact(ptr *[]string) int {
	arr := ptr
	len := 0
	for _, i := range *arr {
		if i != "" {
			len++
		}
	}
	arrTwo := make([]string, len)
	lenTwo := 0
	for _, i := range *arr {
		if i != "" {
			lenTwo++
			arrTwo[lenTwo-1] = i
		}
	}
	*ptr = arrTwo
	return len
}
