package main

import "fmt"

func main() {
	// Create an empty slice with an initial capacity of 5
	x := make([]int, 0, 5)

	// Append values to the slice
	x = append(x, 1, 2, 3, 4)

	// Create two subslices: one from the first two elements, another from the rest
	y := x[:2]
	z := x[2:]

	// Print the capacities of the slices
	fmt.Println(cap(x), len(x), cap(y), len(y), cap(z), len(z)) // Output: 5 5 5

	// Append more elements to the slices
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)

	// Print the slices after appending
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
