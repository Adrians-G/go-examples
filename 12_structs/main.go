package main

import (
	"fmt"
)

func main() {
	person1 := Person{"Samantha", "Smith", "Boston", "f", 25}
	person2 := Person{"James", "White", "Boston", "m", 35}

	person1.marry(person2)
	fmt.Println(person1)
}
