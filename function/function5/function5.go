package main

import "fmt"

func main() {
	// Outer function
	outerFunc := func(start int) func() int {
		count := start // variable in the outer scope
		return func() int {
			count++ // modify the outer variable
			return count
		}
	}

	// Create a closure
	counter := outerFunc(10)

	// Call the closure
	fmt.Println(counter()) // Output: 11
	fmt.Println(counter()) // Output: 12
	fmt.Println(counter()) // Output: 13
}
