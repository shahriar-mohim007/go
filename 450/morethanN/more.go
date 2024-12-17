package main

import (
	"fmt"
)

func moreThanN(arr []int, n int, k int) []int {
	x := n / k
	freq := make(map[int]int)

	for _, num := range arr {
		freq[num]++
	}

	var result []int

	for key, value := range freq {
		if value > x {
			result = append(result, key)
		}
	}

	return result
}

func main() {
	arr := []int{3, 1, 2, 2, 1, 2, 3, 3}
	n := len(arr)
	k := 4
	fmt.Println(moreThanN(arr, n, k))
}
