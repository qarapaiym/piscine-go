package main

import (
	student ".."
	"fmt"
)

type Pilot struct {
	Name     string
	Life     float64
	Age      int
	Aircraft string
}

func main() {
	var donnie student.Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = "1"

	fmt.Println(donnie)
}
