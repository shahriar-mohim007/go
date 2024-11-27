package main

import "fmt"

func main() {
	// Create a map to act as a set, with int as key and struct{} (empty struct) as value.
	intSet := map[int]struct{}{}
	// Sample values, including duplicates.
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	// Populate the set (map) with values.
	for _, v := range vals {
		intSet[v] = struct{}{} // Use an empty struct as the value.
	}
	fmt.Println(intSet)

	// Print the length of vals (input array) and intSet (set with unique values).
	fmt.Println("Length of vals:", len(vals))     // Should print 11
	fmt.Println("Length of intSet:", len(intSet)) // Should print 8 (unique values)

	// Check if certain values exist in the set.
	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set") // Should print this line
	}
	if _, ok := intSet[500]; ok {
		fmt.Println("500 is in the set")
	} else {
		fmt.Println("500 is NOT in the set") // Should print this line
	}

	// Check if 100 is in the set.
	if _, ok := intSet[100]; ok {
		fmt.Println("100 is in the set")
	} else {
		fmt.Println("100 is NOT in the set") // Should print this line
	}
}
