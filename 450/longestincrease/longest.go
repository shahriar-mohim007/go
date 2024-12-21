package main

import "sort"

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sort.Ints(nums)
	maxlength, currlength := 1, 1
	prev := nums[0]

	for i := 1; i < n; i++ {
		if nums[i] == prev+1 {
			currlength++
			maxlength = max(maxlength, currlength)
		} else if nums[i] != prev {
			currlength = 1
		}
		prev = nums[i]
	}
	return maxlength

}
