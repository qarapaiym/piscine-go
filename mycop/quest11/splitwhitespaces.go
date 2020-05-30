package piscine

import "fmt"

func SplitWhiteSpaces(str string) []string {
	len := 0
	for i := range str {
		if string(str[i]) == " " && string(str[i+1]) != " " {
			len++
		}
	}

	arr := make([]string, len+1)

	lenArr := 0
	strArr := ""
	for i, c := range str {
		if string(c) != " " {
			strArr += string(c)
		}
			arr[lenArr] = strArr
		if string(c) == " " && string(str[i+1]) != " " && strArr != "" {
			lenArr++
			strArr = ""
		}
	}
	return arr
}
