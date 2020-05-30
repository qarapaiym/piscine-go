package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	length := 0
	for i := range args {
		length = i + 1
	}

	if length != 3 {
		return
	}

	num1_str := args[0]

	for i, s := range num1_str {
		if (s >= '0' && s <= '9') || (i == 0 && s == '-') {
			continue
		} else {
			os.Stdout.WriteString("0\n")
			return
		}
	}
	num2_str := args[2]
	for i, s := range num2_str {
		if (s >= '0' && s <= '9') || (i == 0 && s == '-') {
			continue
		} else {
			os.Stdout.WriteString("0\n")
			return
		}
	}

	num1 := Atoi(num1_str)
	num2 := Atoi(num2_str)
	op_str := args[1]
	if op_str == "main.go" {
		op_str = "*"
	}
	len_op := 0
	for i := range op_str {
		len_op = i + 1
	}
	if len_op != 1 {
		os.Stdout.WriteString("HERE")

		os.Stdout.WriteString("0\n")
		return
	}

	num := 0
	op := op_str[0]

	if op == '+' {
		num = num1 + num2
	} else if op == '-' {
		num = num1 - num2
	} else if op == '*' {
		num = num1 * num2
	} else if op == '/' {
		if num2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		num = num1 / num2
	} else if op == '%' {
		if num2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		num = num1 % num2
	} else {
		os.Stdout.WriteString("0\n")
		return
	}

	final := ConvertInt64ToString(num)
	if (num > -9223372036854775808 && num < 9223372036854775807) && final != "-" && (num1 > -9223372036854775807 && num1 < 9223372036854775807) && (num2 > -9223372036854775808 && num2 < 9223372036854775807) {
		os.Stdout.WriteString(final)
		os.Stdout.WriteString("\n")
	} else {
		os.Stdout.WriteString("0\n")
	}
}

func Atoi(s string) int {
	x := 0
	z := 1
	for i, n := range s {
		y := 0
		if n < '0' || n > '9' {
			if n == '-' || n == '+' {
				if i != 0 {
					return 0
				}
				if n == '-' {
					z = -1
				}
			} else {
				return 0
			}
		}
		for i := '1'; i <= n; i++ {
			y++
		}
		x = x*10 + y*z
	}
	return x
}

func ConvertInt64ToString(n int) string {
	str := ""
	str1 := ""
	if n == 0 {
		str1 = "0"
	} else if n < 0 {
		n = -n
		for n >= 1 {
			str = str + string(n%10+48)
			n = n / 10
		}
		str = str + "-"
	} else {
		for n >= 1 {
			str = str + string(n%10+48)
			n = n / 10
		}
	}
	l := 0
	for range str {
		l++
	}
	for i := l - 1; i >= 0; i-- {
		str1 = str1 + string(str[i])
	}
	if str1 != "-" {
		return str1
	}
	return "0"
}
