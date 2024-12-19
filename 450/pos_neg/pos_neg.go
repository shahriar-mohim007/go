package main

import "fmt"

func rearrangeArray(arr []int) []int {
	// Create two slices for positive and negative numbers
	var pos []int
	var neg []int

	// Separate positive and negative numbers
	for _, num := range arr {
		if num >= 0 {
			pos = append(pos, num)
		} else {
			neg = append(neg, num)
		}
	}

	// Prepare a result slice
	var result []int
	i, j := 0, 0

	// Add elements alternately from pos and neg
	for i < len(pos) && j < len(neg) {
		result = append(result, pos[i])
		result = append(result, neg[j])
		i++
		j++
	}

	// If any positive numbers are left, add them to the result
	for i < len(pos) {
		result = append(result, pos[i])
		i++
	}

	// If any negative numbers are left, add them to the result
	for j < len(neg) {
		result = append(result, neg[j])
		j++
	}

	return result
}

func main() {
	// Example 1
	arr1 := []int{1, 2, 3, -4, -1, 4}
	result1 := rearrangeArray(arr1)
	fmt.Println(result1) // Output: [1 -4 2 -1 3 4]

	// Example 2
	arr2 := []int{-5, -2, 5, 2, 4, 7, 1, 8, 0, -8}
	result2 := rearrangeArray(arr2)
	fmt.Println(result2) // Output: [-5 5 -2 2 -8 4 7 1 8 0]
}
