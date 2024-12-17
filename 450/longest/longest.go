package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Sort the array
	sort.Ints(nums)

	maxLength := 1
	curLength := 1
	prevElement := nums[0]

	// Iterate through the array
	for i := 1; i < len(nums); i++ {
		if nums[i] == prevElement+1 {
			curLength++
			maxLength = max(maxLength, curLength)
		} else if nums[i] != prevElement {
			curLength = 1
		}
		prevElement = nums[i]
	}

	return maxLength
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums)) // Output: 4 (sequence: [1, 2, 3, 4])
}
