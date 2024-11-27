package main

import (
	"fmt"
	"slices"
)

func main() {
	// Example slice
	t := []int{1, 2, 3, 4, 5}
	t2 := []int{1, 2, 3, 4, 5}
	t3 := []int{5, 4, 3, 2, 1}

	// 1. Equal: Compare two slices
	if slices.Equal(t, t2) {
		fmt.Println("t is equal to t2")
	} else {
		fmt.Println("t is not equal to t2")
	}

	if slices.Equal(t, t3) {
		fmt.Println("t is equal to t3")
	} else {
		fmt.Println("t is not equal to t3")
	}

	// 2. Clone: Create a copy of a slice
	clone := slices.Clone(t)
	fmt.Println("Clone of t:", clone)

	// 3. Contains: Check if a value exists in the slice
	fmt.Println("Does t contain 3?", slices.Contains(t, 3))   // true
	fmt.Println("Does t contain 10?", slices.Contains(t, 10)) // false

	// 4. Index: Find the index of a value
	fmt.Println("Index of 3 in t:", slices.Index(t, 3))   // 2
	fmt.Println("Index of 10 in t:", slices.Index(t, 10)) // -1

	// 5. Delete: Remove elements by index
	deleted := slices.Delete(t, 1, 3) // Removes elements at index 1 and 2
	fmt.Println("After deleting elements at index 1 and 2:", deleted)

	// 6. Replace: Replace a range of elements with another slice
	//replaced := slices.Replace(t, 1, 3, []int{7, 8})
	//fmt.Println("After replacing elements at index 1 and 2 with [7, 8]:", replaced)

	// 7. Sort: Sort the slice
	tUnsorted := []int{5, 3, 1, 4, 2}
	slices.Sort(tUnsorted)
	fmt.Println("Sorted slice:", tUnsorted)

	// 8. SortFunc: Sort with a custom comparator (descending order)
	//slices.SortFunc(tUnsorted, func(a, b int) bool { return a > b })
	//fmt.Println("Sorted in descending order:", tUnsorted)

	// 9. BinarySearch: Find an element in a sorted slice
	index, found := slices.BinarySearch(t, 3)
	fmt.Printf("Binary search for 3: index = %d, found = %t\n", index, found)

	// 10. Insert: Insert elements at a specific index
	inserted := slices.Insert(t, 2, 99)
	fmt.Println("After inserting 99 at index 2:", inserted)

	// 11. Compact: Remove adjacent duplicates
	duplicates := []int{1, 2, 2, 3, 3, 3, 4}
	compact := slices.Compact(duplicates)
	fmt.Println("After compacting duplicates:", compact)

	// 12. CompactFunc: Compact with a custom function
	compactCustom := slices.CompactFunc(duplicates, func(a, b int) bool { return a == b })
	fmt.Println("After compacting with a custom function:", compactCustom)
}
