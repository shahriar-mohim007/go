package main

import (
	"fmt"
	"sort"
)

func find3Numbers(nums []int, n int, x int) []int {
	sort.Ints(nums) // Sort the array

	for i := 0; i < n-2; i++ {
		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == x {
				return []int{nums[i], nums[left], nums[right]}
			} else if sum < x {
				left++
			} else {
				right--
			}
		}
	}

	return []int{}
}

func main() {
	nums := []int{1, 4, 45, 6, 10, 8}
	n := len(nums)
	x := 22

	result := find3Numbers(nums, n, x)
	if len(result) > 0 {
		fmt.Println("Triplet found:", result)
	} else {
		fmt.Println("No triplet found")
	}
}
