package main

import "fmt"

func main() {
	ids := []int{33, 39, 99, 121}

	// Loop trough
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0

	for _, id := range ids {
		sum += id
	}

	// Range with map
	emails := map[string]string{"Bob": "bog@gmail.com", "x": "y"}

	for k, v := range emails {
		fmt.Printf("Key: %s // Value: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: ", k)
	}
}
