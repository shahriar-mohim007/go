package main

import "fmt"

// Original function: Does not update the original pointer
func failedUpdate(g *int) {
	x := 11
	g = &x // Reassigns the local copy of g to point to x
}

// Corrected function: Updates the original pointer by using **int
func successfulUpdate(g **int) {
	x := 10
	*g = &x // Updates the original pointer by dereferencing g
}

func ValueUpdate(g *int) {
	*g = 111
}

func main() {
	// Demonstrating failedUpdate
	var f *int                             // f is nil
	fmt.Println("Before failedUpdate:", f) // Prints: nil
	failedUpdate(f)
	fmt.Println("After failedUpdate:", f) // Prints: nil

	x := 10
	ValueUpdate(&x)
	fmt.Println("after successfulUpdate:", x)
	// Demonstrating successfulUpdate
	var f2 *int                                 // f2 is nil
	fmt.Println("Before successfulUpdate:", f2) // Prints: nil
	successfulUpdate(&f2)
	fmt.Println("After successfulUpdate:", f2)    // Prints: memory address
	fmt.Println("Dereferenced value of f2:", *f2) // Prints: 10
}
