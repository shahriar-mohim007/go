package main

import "fmt"

func CustomAppend(slice []int, elements ...int) []int {
	// If the slice has enough capacity, directly append the elements
	if len(slice)+len(elements) <= cap(slice) {
		slice = append(slice, elements...)
		return slice
	}

	// If the slice doesn't have enough capacity, we need to create a larger slice
	// Allocate a new slice with double the current capacity (or enough to hold both old and new elements)
	newSlice := make([]int, len(slice), (len(slice)+len(elements))*2)

	// Copy the old slice into the new slice
	copy(newSlice, slice)

	// Append the new elements
	newSlice = append(newSlice, elements...)

	return newSlice
}
func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, 30, 50, 60)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println(cap(x), cap(y))
}
