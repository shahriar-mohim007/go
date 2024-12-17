package main

import "fmt"

func main() {
	// Create a map to act as a set, with int as key and bool as value.
	intSet := map[int]bool{}
	// Sample values, including duplicates.
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	// Populate the set (map) with values.
	for _, v := range vals {
		intSet[v] = true
	}

	// Print the length of vals (input array) and intSet (set with unique values).
	fmt.Println("Length of vals:", len(vals))     // Should print 11
	fmt.Println("Length of intSet:", len(intSet)) // Should print 8 (unique values)

	// Check if certain values exist in the set.
	fmt.Println("set", intSet)
	fmt.Println("Is 5 in the set?", intSet[5])     // Should print true
	fmt.Println("Is 500 in the set?", intSet[500]) // Should print false

	// Check if 100 is in the set.
	if intSet[100] {
		fmt.Println("100 is in the set")
	} else {
		fmt.Println("100 is NOT in the set") // Should print this line
	}
}
