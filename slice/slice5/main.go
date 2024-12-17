package main

import "fmt"

// contains checks if a slice contains a specific value
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// indexOf returns the index of a value in a slice, or -1 if not found
func indexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// remove removes an element at a specified index from the slice
func remove(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice // Return the original slice if index is invalid
	}
	return append(slice[:index], slice[index+1:]...)
}

// reverse reverses the elements in a slice
func reverse(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func CustomAppend(slice []int, elements ...int) []int {
	// Check if the slice has enough capacity
	if len(slice)+len(elements) <= cap(slice) {
		// If enough capacity, simply add new elements to the end
		for _, elem := range elements {
			slice = appendToSlice(slice, elem)
		}
		return slice
	}

	// If not enough capacity, create a new larger slice
	// Double the current capacity (or enough to hold the current + new elements)
	newSlice := make([]int, len(slice), (len(slice)+len(elements))+1)

	// Copy the existing elements into the new slice
	for i := 0; i < len(slice); i++ {
		newSlice[i] = slice[i]
	}

	// Append new elements to the new slice
	for _, elem := range elements {
		newSlice = appendToSlice(newSlice, elem)
	}

	return newSlice
}

// appendToSlice manually appends an element to the slice by increasing its length
func appendToSlice(slice []int, element int) []int {
	// Increase the length of the slice by 1
	slice = slice[:len(slice)+1]
	// Add the new element to the last position
	slice[len(slice)-1] = element
	return slice
}

// filter returns a new slice containing elements that satisfy a predicate
func filter(slice []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Test slice
	slice := []int{10, 20, 30, 40, 50}

	// contains
	fmt.Println("Slice contains 30?", contains(slice, 30)) // true
	fmt.Println("Slice contains 60?", contains(slice, 60)) // false

	// indexOf
	fmt.Println("Index of 30:", indexOf(slice, 30)) // 2
	fmt.Println("Index of 60:", indexOf(slice, 60)) // -1

	// remove
	fmt.Println("Slice after removing index 2:", remove(slice, 2)) // [10 20 40 50]
	fmt.Println("Slice after removing index 5:", remove(slice, 5)) // No change

	// reverse
	reversed := make([]int, len(slice)) // Create a copy to reverse
	copy(reversed, slice)
	fmt.Println("Reversed slice:", reverse(reversed)) // [50 40 30 20]

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Define a predicate: we want only even numbers
	isEven := func(n int) bool {
		return n%2 == 0
	}

	// Call filter with the slice and predicate
	evenNumbers := filter(numbers, isEven)

	// Output the result
	fmt.Println("Even numbers:", evenNumbers)
	X := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Y := X[:]
	fmt.Println(reverse(Y))
	fmt.Println(X)

	x := []int{1, 2, 3, 4}
	fmt.Println(x[0])
	y := x[:2]
	z := x[1:]
	d := x[1:2]
	e := x[:]

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	//first index inclusive and last index exclusive

	arr := [5]int{1, 2, 3, 4, 5}

	// Create a slice from the array
	Slice := arr[:]
	Slice2 := arr[:]

	// Print the slice and its underlying array
	fmt.Println("Array:", arr)   // Original array: [1, 2, 3, 4, 5]
	fmt.Println("Slice:", Slice) // Slice: [2, 3, 4]

	// Modify an element in the slice
	Slice[1] = 99
	Slice2[2] = 100
	Slice2 = append(Slice2, 100)
	// Print the array and slice after modification
	fmt.Println("After modifying the slice:")
	fmt.Println("Array:", arr)   // Array is affected: [1, 2, 99, 4, 5]
	fmt.Println("Slice:", Slice) // Slice reflects the change: [2, 99, 4]
	fmt.Println("Slice2:", Slice2)

	slice = CustomAppend(slice, 4, 5, 6)
	fmt.Println("After first append:", slice)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))

	// Append more elements, causing reallocation
	slice = CustomAppend(slice, 7, 8, 9, 10)
	fmt.Println("After second append:", slice)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))

}

//Slices share storage sometimes
//When you take a slice from a slice, you are not making a copy of the data.
//Instead, you now have two variables that are sharing memory. This means
//that changes to an element in a slice affect all slices that share that element.
//Let’s see what happens when we change values.
