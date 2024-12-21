package main

func containsDuplicate(nums []int) bool {
	duplicates := make(map[int]bool)
	for _, n := range nums {
		if duplicates[n] {
			return true
		}
		duplicates[n] = true
	}
	return false
}
