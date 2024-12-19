package main

import "fmt"

// Function to count inversions using brute force
func countInversions(arr []int) int {
	count := 0
	n := len(arr)

	// Iterate through each pair (i, j) where i < j
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				count++ // Increment count if arr[i] > arr[j]
			}
		}
	}
	return count
}

func main() {
	arr1 := []int{2, 4, 1, 3, 5}
	fmt.Println("Inversion Count:", countInversions(arr1)) // Output: 3

	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Inversion Count:", countInversions(arr2)) // Output: 0

	arr3 := []int{5, 4, 3, 2, 1}
	fmt.Println("Inversion Count:", countInversions(arr3)) // Output: 10
}
