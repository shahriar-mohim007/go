package main

import (
	"fmt"
	"math"
)

func trapWater(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	// Create arrays to store the maximum heights to the left and right
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	// Fill leftMax
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = int(math.Max(float64(leftMax[i-1]), float64(height[i])))
	}

	// Fill rightMax
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = int(math.Max(float64(rightMax[i+1]), float64(height[i])))
	}

	// Calculate trapped water
	water := 0
	for i := 0; i < n; i++ {
		water += int(math.Max(0, math.Min(float64(leftMax[i]), float64(rightMax[i]))-float64(height[i])))
	}

	return water
}

func main() {
	arr := []int{0, 0, 0, 2, 5, 2}
	fmt.Println("Water trapped:", trapWater(arr))
}
