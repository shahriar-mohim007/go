package main

import (
	"fmt"
)

func threeWayPartition(arr []int, a, b int) bool {
	low := 0
	mid := 0
	high := len(arr) - 1

	for mid <= high {
		if arr[mid] < a {
			// Swap arr[low] and arr[mid], increment both pointers
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] >= a && arr[mid] <= b {
			// Element in range, move mid pointer
			mid++
		} else {
			// Swap arr[mid] and arr[high], decrement high pointer
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}

	return true
}

func main() {
	arr := []int{1, 2, 3, 3, 4}
	a := 1
	b := 2

	fmt.Println("Original array:", arr)
	if threeWayPartition(arr, a, b) {
		fmt.Println("Partitioned array:", arr)
	} else {
		fmt.Println("Partitioning failed.")
	}
}
