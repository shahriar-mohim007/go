package main

import (
	"fmt"
	"math"
	"sort"
)

func findMinDiff(arr []int, m int) int {
	n := len(arr)

	// If there are fewer packets than students, return -1 (invalid input)
	if n < m {
		return -1
	}

	// Sort the array
	sort.Ints(arr)

	// Initialize the minimum difference to a very large value
	minDiff := math.MaxInt

	// Iterate over the array to find the minimum difference
	for i := 0; i <= n-m; i++ {
		currentDiff := arr[i+m-1] - arr[i] // Difference between max and min in the window
		if currentDiff < minDiff {
			minDiff = currentDiff
		}
	}

	return minDiff
}

func main() {
	arr := []int{3, 4, 1, 9, 56, 7, 9, 12}
	m := 5

	result := findMinDiff(arr, m)
	if result != -1 {
		fmt.Println("The minimum difference is:", result)
	} else {
		fmt.Println("Not enough packets for the students.")
	}
}
