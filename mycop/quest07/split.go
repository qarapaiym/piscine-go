package piscine

func Split(str, charset string) []string {
	length := 0
	for i := range str {
		length = i + 1

	}
	lenSplit := 0
	for i := range charset {
		lenSplit = i + 1
	}
	for i := 0; i < lenSplit; i++ {
		str += " "
	}
	prev := false
	len := 0
	for i := 0; i < length; i++ {
		if (str[i:i+lenSplit] == charset) && !prev {
			prev = true
			len++
		} else {
			prev = false
		}
	}
	len++
	arr := make([]string, len)

	stringa := ""
	count := 0
	for i := 0; i < length; i++ {
		if str[i:i+lenSplit] == charset {
			l := 0
			for i := range stringa {
				l = i + 1
			}
			if l == 0 {
				continue
			}
			arr[count] = stringa
			count++
			stringa = ""
			i = i + lenSplit - 1
			continue
		}
		stringa += string(str[i])
	}

	l := 0
	for i := range stringa {
		l = i + 1
	}
	if l != 0 {
		arr[count] = stringa
	}
	return arr
}
