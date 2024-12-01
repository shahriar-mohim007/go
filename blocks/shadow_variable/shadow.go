package main

import "fmt"

func main() {
	x := 10 // Declare x in the outer scope
	if x > 5 {
		fmt.Println("Outer x:", x) // Prints 10
		x, y := 5, 20
		fmt.Println("Inner", x, y) // Declare a new x in the inner scope (shadowing the outer x)
	}
	fmt.Println("Outer x after if block:", x) // Prints 10
}
