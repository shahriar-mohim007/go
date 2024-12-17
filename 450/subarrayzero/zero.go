package main

import "fmt"

func hasZeroSumSubarray(arr []int) bool {
	prefixSum := 0
	prefixSet := make(map[int]bool)

	for _, num := range arr {
		prefixSum += num

		// Check if prefixSum is 0, or if it already exists in the set
		if prefixSum == 0 || prefixSet[prefixSum] {
			return true
		}

		// Add prefixSum to the set
		prefixSet[prefixSum] = true
	}

	return false
}

func main() {
	arr := []int{4, 2, -3, 1, 6}
	fmt.Println(hasZeroSumSubarray(arr)) // Output: true

	arr = []int{4, 2, 0, 1, 6}
	fmt.Println(hasZeroSumSubarray(arr)) // Output: true (0 alone is a subarray with sum 0)

	arr = []int{-3, 2, 3, 1, 6}
	fmt.Println(hasZeroSumSubarray(arr)) // Output: false
}
