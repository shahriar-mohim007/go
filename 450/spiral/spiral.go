package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)
	rowbegin := 0
	rowend := len(matrix) - 1 // last row index
	colbegin := 0
	colend := len(matrix[0]) - 1 // last column index

	for rowbegin <= rowend && colbegin <= colend {
		// Traverse from left to right
		for i := colbegin; i <= colend; i++ {
			result = append(result, matrix[rowbegin][i])
		}
		rowbegin++

		// Traverse from top to bottom
		for i := rowbegin; i <= rowend; i++ {
			result = append(result, matrix[i][colend])
		}
		colend--

		// Traverse from right to left (only if we have more than one row left)
		if rowbegin <= rowend {
			for i := colend; i >= colbegin; i-- {
				result = append(result, matrix[rowend][i])
			}
			rowend--
		}

		// Traverse from bottom to top (only if we have more than one column left)
		if colbegin <= colend {
			for i := rowend; i >= rowbegin; i-- {
				result = append(result, matrix[i][colbegin])
			}
			colbegin++
		}
	}

	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result := spiralOrder(matrix)
	fmt.Println(result) // Output: [1 2 3 6 9 8 7 4 5]
}
