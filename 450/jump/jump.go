package main

import "fmt"

func canJump(nums []int) bool {
	length := len(nums)
	dp := make([]int, length)
	dp[0] = nums[0]
	for i := 1; i < length; i++ {
		if dp[i-1] < i {
			return false
		}
		dp[i] = max(i+nums[i], dp[i-1])
		if dp[i] >= length-1 {
			return true
		}
	}
	return dp[length-1] >= length-1
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
