package main

import "fmt"

func maxProfit(prices []int) int {
	left, right := 0, 1
	ans := 0
	for right < len(prices) {
		if prices[right] > prices[left] {
			cur := prices[right] - prices[left]
			ans = max(ans, cur)
		} else {
			left = right
		}
		right++
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
