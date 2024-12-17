package main

import (
	"fmt"
)

func maxProduct(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	// Initialize max and min products to the first element
	maxProduct := arr[0]
	minProduct := arr[0]
	result := arr[0]

	// Iterate through the array from the second element
	for i := 1; i < len(arr); i++ {
		// If the current number is negative, swap maxProduct and minProduct
		if arr[i] < 0 {
			maxProduct, minProduct = minProduct, maxProduct
		}

		// Update maxProduct and minProduct
		maxProduct = max(arr[i], arr[i]*maxProduct)
		minProduct = min(arr[i], arr[i]*minProduct)

		// Update the result with the maximum product found so far
		result = max(result, maxProduct)
	}

	return result
}

func main() {
	arr := []int{-2, 6, -3, -10, 0, 2}
	fmt.Println(maxProduct(arr)) // Output: 180
}
