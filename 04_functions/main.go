package main

import "fmt"

func test(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(test("World"))
}
