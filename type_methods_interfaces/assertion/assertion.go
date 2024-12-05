package main

import "fmt"

func main() {
	var i interface{} = 42 // Interface holding an int

	// Type assertion
	value, ok := i.(int)
	if ok {
		fmt.Println("Value:", value) // Output: Value: 42
	} else {
		fmt.Println("Type assertion failed")
	}

	// Unsafe assertion
	// value := i.(string) // This would panic because i doesn't hold a string
}
