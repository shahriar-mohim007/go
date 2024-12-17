package main

import "fmt"

// Function to modify an array using a pointer
func zeroArray(arr *[5]int) {
	for i := range arr {
		arr[i] = 0
	}
}

// Function to modify a slice
func zeroSlice(slice []int) {
	for i := range slice {
		slice[i] = 0
	}
}

func main() {
	// Using an array
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", array)

	// Passing a pointer to the array to the function
	zeroArray(&array)
	fmt.Println("Modified array:", array)

	// Using a slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", slice)

	// Passing the slice to the function
	zeroSlice(slice)
	fmt.Println("Modified slice:", slice)

	// Demonstrating slice flexibility
	slice = append(slice, 6, 7, 8)
	fmt.Println("Appended slice:", slice)

	// Slicing from an array
	newSlice := array[1:4] // Creates a slice backed by the array
	fmt.Println("Slice from array:", newSlice)

	// Modifying the slice also modifies the underlying array
	zeroSlice(newSlice)
	fmt.Println("Modified slice from array:", newSlice)
	fmt.Println("Underlying array after slice modification:", array)
}
