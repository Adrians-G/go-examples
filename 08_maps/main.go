package main

import "fmt"

func main() {
	// // Define map [key type]value type
	// emails := make(map[string]string)

	// // Assign
	// emails["Bob"] = "bob@gmail.com"
	// emails["Def"] = "default@gmail.com"

	// Short hand
	emails := map[string]string{"Bob": "bog@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")
	fmt.Println(emails)
}
