package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	var name = "Hello"
	var age = 5
	var isCool = true

	const something = "123"
	//	something = "213" -> error

	// short-hand
	namee := "Brad"
	size := 1.3

	// double assign

	sun, moon := "Sun", "Moon"

	fmt.Println(name, age, isCool, namee, size)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
	fmt.Println(sun, moon)
}
