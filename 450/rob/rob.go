package main

import "fmt"

func robLinear(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n+1)
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	maxRob1 := robLinear(nums[1:])
	maxRob2 := robLinear(nums[:n-1])
	return max(maxRob1, maxRob2)
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 4, 5}))
}
