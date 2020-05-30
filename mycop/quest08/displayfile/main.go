package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	names := os.Args[1:]
	len := 0
	for range names {
		len++
	}
	name := ""
	if len == 1 {
		name = names[0]
	}
	if len > 1 {
		fmt.Println("Too many arguments")
	} else if name == "quest8.txt" || name == "./student/displayfile/quest8.txt" {
		b, _ := ioutil.ReadFile(name)
		fmt.Println(string(b))
	} else {
		fmt.Println("File name missing")
	}
}
