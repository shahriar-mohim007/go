package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxarea := 0
	for left < right {
		width := right - left
		current := min(height[left], height[right]) * width
		maxarea = max(maxarea, current)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxarea
}
