package main

import (
	"fmt"
)

// minSwaps finds the minimum number of swaps required to group all elements ≤ k together
func minSwaps(arr []int, k int) int {
	// Step 1: Count elements ≤ k
	count := 0
	for _, num := range arr {
		if num <= k {
			count++
		}
	}

	// Step 2: Count elements > k in the first window of size `count`
	bad := 0
	for i := 0; i < count; i++ {
		if arr[i] > k {
			bad++
		}
	}

	// Step 3: Use sliding window to find the minimum bad elements
	minBad := bad
	for i := 0; i < len(arr)-count; i++ {
		// Remove the first element of the current window
		if arr[i] > k {
			bad--
		}
		// Add the next element into the window
		if arr[i+count] > k {
			bad++
		}
		// Update the minimum bad count
		if bad < minBad {
			minBad = bad
		}
	}

	// Minimum bad elements = minimum swaps required
	return minBad
}

func main() {
	arr := []int{2, 1, 5, 6, 3}
	k := 3
	fmt.Println("Minimum swaps required:", minSwaps(arr, k))
}

//func main() {
//	var n int
//	fmt.Println("Enter the number of elements in the array:")
//	fmt.Scan(&n)
//
//	nums := make([]int, n)
//	fmt.Println("Enter the elements of the array:")
//	for i := 0; i < n; i++ {
//		fmt.Scan(&nums[i])
//	}
//
//	fmt.Println("Enter the k:")
//	fmt.Scan(&n)
//}
