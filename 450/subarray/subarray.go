package main

import "fmt"

func maxSubArray(nums []int) int {
	sum, ans := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		ans = max(ans, sum)
		sum = max(sum, 0)
	}
	return ans
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
