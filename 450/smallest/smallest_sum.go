package main

import (
	"fmt"
	"math"
)

func smallestSum(nums []int, x int) int {
	n := len(nums)
	minLength := math.MaxInt // Initialize with a large value
	currentSum := 0
	start := 0

	for end := 0; end < n; end++ {
		// Add the current element to the sum
		currentSum += nums[end]

		// Shrink the window as long as the sum is greater than x
		for currentSum > x {
			// Update the minimum length
			minLength = min(minLength, end-start+1)

			// Shrink the window from the left
			currentSum -= nums[start]
			start++
		}
	}

	// If no subarray found, return 0
	if minLength == math.MaxInt {
		return 0
	}

	return minLength
}

func main() {
	arr := []int{1, 4, 45, 6, 0, 19}
	x := 51
	result := smallestSum(arr, x)
	fmt.Println("Minimum length of subarray with sum greater than", x, "is:", result)
}
