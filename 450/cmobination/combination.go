package main

import "fmt"

func combine(index int, result *[][]int, path []int, candidates []int, target int, currentSum int) {
	if currentSum == target {
		*result = append(*result, append([]int{}, path...))
		return
	}
	if currentSum > target {
		return
	}

	for j := index; j < len(candidates); j++ {
		// Add the current candidate to path and update currentSum
		path = append(path, candidates[j])
		combine(j, result, path, candidates, target, currentSum+candidates[j])
		// Backtrack by removing the last element and reverting currentSum
		path = path[:len(path)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	combine(0, &result, []int{}, candidates, target, 0)
	return result
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := combinationSum(candidates, target)
	fmt.Println(result) // Output: [[2 2 3] [7]]
}
