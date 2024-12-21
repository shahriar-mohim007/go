package main

func twoSum(nums []int, target int) []int {
	indexmap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := indexmap[complement]; ok {
			return []int{i, j}
		}
		indexmap[num] = i
	}
	return nil
}
