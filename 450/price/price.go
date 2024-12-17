package main

import (
	"fmt"
)

func sellingShare(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	// Step 1: Create a profit array
	profit := make([]int, n)

	// Step 2: Backward pass (right to left)
	maxPrice := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		// Update maxPrice seen so far
		maxPrice = max(maxPrice, nums[i])
		// Update profit[i] for one transaction ending after day i
		profit[i] = max(profit[i+1], maxPrice-nums[i])
	}

	// Step 3: Forward pass (left to right)
	minPrice := nums[0]
	for i := 1; i < n; i++ {
		// Update minPrice seen so far
		minPrice = min(minPrice, nums[i])
		// Combine profits from two transactions
		profit[i] = max(profit[i-1], profit[i]+(nums[i]-minPrice))
	}

	// Final maximum profit
	return profit[n-1]
}

func main() {
	price := []int{10, 22, 5, 75, 65, 80}
	fmt.Println("Maximum Profit:", sellingShare(price))
}
