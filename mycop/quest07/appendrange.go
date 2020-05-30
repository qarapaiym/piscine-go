package piscine

func AppendRange(min, max int) []int {
	var args []int
	if min >= max {
		return args
	}

	for i := min; i < max; i++ {
		args = append(args, i)
	}
	return args
}
