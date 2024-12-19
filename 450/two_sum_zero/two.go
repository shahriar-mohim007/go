package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int) [][]int {

	seen := make(map[int]bool)
	pairSet := make(map[[2]int]bool)
	var result [][]int

	for _, num := range nums {
		if seen[num] {
			pair := [2]int{min(num, -num), max(num, -num)}
			pairSet[pair] = true
		}
		seen[-num] = true
	}

	for pair := range pairSet {
		result = append(result, []int{pair[0], pair[1]})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(twoSum(nums))
	nums = []int{6, 1, -6, 8, 0, 4, -9, -1, -10, -5, 6}
	fmt.Println(twoSum(nums))
}
