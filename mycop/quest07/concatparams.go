package piscine

func ConcatParams(args []string) string {
	str := ""
	first := false
	for i := range args {
		if first {
			str += "\n"
		}
		str += args[i]
		first = true
	}
	return str
}
