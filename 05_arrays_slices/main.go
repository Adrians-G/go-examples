package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	// Arrays
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// fmt.Println(fruitArr)

	// Declare and assign
	// shortHand := [2]int{10, 50}

	// fmt.Println(shortHand)

	// Slices (Dynamic length)
	fruitSlice := []string{"abc", "efg", "xxx", "yyy"}

	spew.Dump(fruitSlice[0:2])
	spew.Dump(len(fruitSlice))
}
