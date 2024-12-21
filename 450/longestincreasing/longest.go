package main

import "fmt"

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	ans := 1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				ans = max(ans, dp[i])
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLIS([]int{1, 3, 5, 6}))
}
