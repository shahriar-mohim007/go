package main

import "fmt"

func main() {
	// Create a slice x with capacity 5
	x := make([]int, 0, 5)

	// Append elements to x
	x = append(x, 1, 2, 3, 4)

	// Create full slice expressions
	y := x[:2:2]  // y starts at 0, ends at index 2, with capacity 2
	z := x[2:4:4] // z starts at index 2, ends at index 4, with capacity 2

	// Print the capacities of x, y, and z
	fmt.Println("Capacity of x, y, z:", cap(x), cap(y), cap(z))

	// Append more elements to y, x, and z
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)

	// Print the final values of x, y, and z
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

//full slice expressions
//slice[startIndex:endIndex:maxCapacity]
