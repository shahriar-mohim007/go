package main

import (
	"fmt"
	"strconv"
)

// palinCheck checks if a given number is a palindrome
func palinCheck(num int) bool {
	numStr := strconv.Itoa(num)
	left := 0
	right := len(numStr) - 1

	for left <= right {
		if numStr[left] == numStr[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

// PalinArray checks if all numbers in the array are palindromes
func PalinArray(nums []int) bool {
	for _, num := range nums {
		if !palinCheck(num) {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Println("Enter the number of elements in the array:")
	fmt.Scan(&n)

	nums := make([]int, n)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println("Input Array:", nums)
	if PalinArray(nums) {
		fmt.Println("All numbers in the array are palindromes.")
	} else {
		fmt.Println("Not all numbers in the array are palindromes.")
	}
}
