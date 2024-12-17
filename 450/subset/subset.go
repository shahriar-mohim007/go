package main

import "fmt"

func isSubset(a, b []int) string {
	// Create frequency maps for a and b
	countA := make(map[int]int)
	countB := make(map[int]int)

	// Count elements in a
	for _, num := range a {
		countA[num]++
	}

	// Count elements in b
	for _, num := range b {
		countB[num]++
	}

	// Check if b is a subset of a
	for key, count := range countB {
		if countA[key] < count {
			return "No"
		}
	}

	return "Yes"
}

func main() {
	a := []int{11, 7, 1, 13, 21, 3, 7, 3}
	b := []int{11, 3, 7, 1, 7}

	result := isSubset(a, b)
	fmt.Println(result) // Output: Yes
}
