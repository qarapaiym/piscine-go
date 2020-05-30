package main

func main(str string) int {
	aChangable := []rune(str)
	len := 0
	for i := range aChangable {
		len = i
	}
	len++
	return len
}
