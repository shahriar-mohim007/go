package main

import (
	"fmt"
	"sort"
)

func main() {
	// Initialize a slice
	slice := []int{30, 10, 20, 50, 40}

	// Length and capacity
	fmt.Println("Slice:", slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// Append elements
	slice = append(slice, 60, 70)
	fmt.Println("After appending:", slice)

	// Sorting (uses the sort package)
	sort.Ints(slice)
	fmt.Println("Sorted slice:", slice)

	// Access elements by index
	fmt.Println("First element:", slice[0])
	fmt.Println("Last element:", slice[len(slice)-1])

	// Sub-slice
	subSlice := slice[1:4]
	fmt.Println("Sub-slice (1:4):", subSlice)

	// Copying a slice
	copiedSlice := make([]int, len(slice))
	copy(copiedSlice, slice)
	fmt.Println("Copied slice:", copiedSlice, slice)

	// Iterating using range
	fmt.Println("Iterating over slice:")
	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// Removing an element (manual, using slicing)
	indexToRemove := 2 // Remove element at index 2
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	fmt.Println("After removing element at index 2:", slice)
}
