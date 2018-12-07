package main

import "fmt"

func printStuff(from string) {
	for i := 0; i <= 9999999; i++ {
		for i := 0; i <= 9999999; i++ {
			if i == 9999999 {
				fmt.Println(i)
				return
			}
		}
	}
}

func main() {
	printStuff("1337")
}
