package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}
	postfix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}
	return result
}
