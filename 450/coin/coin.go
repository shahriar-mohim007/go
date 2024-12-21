package main

import "fmt"

func coin(index, amount, n int, coins []int, dp [][]int) int {
	if amount < 0 || index == n {
		return int(1e9) // Return a large value for invalid states
	}
	if amount == 0 {
		return 0 // If amount is 0, no more coins are needed
	}
	if dp[index][amount] != -1 {
		return dp[index][amount] // Return already computed result
	}

	// Option 1: Take the current coin
	res1 := 1 + coin(index, amount-coins[index], n, coins, dp)
	// Option 2: Skip the current coin
	res2 := coin(index+1, amount, n, coins, dp)

	// Take the minimum of both options
	dp[index][amount] = min(res1, res2)
	return dp[index][amount]
}

func coinChange(coins []int, amount int) int {
	n := len(coins)
	// Initialize the memoization table
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	// Start the recursive computation
	result := coin(0, amount, n, coins, dp)
	if result == int(1e9) {
		return -1 // If no solution, return -1
	}
	return result
}
func main() {
	fmt.Println(coinChange([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))
}
