package main

import "fmt"

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	// Create two boolean slices to track rows and columns that need to be zeroed
	rows := make([]bool, m)
	cols := make([]bool, n)

	// First pass: identify rows and columns to be zeroed
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	// Second pass: zero out the identified rows
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}

	setZeroes(matrix)

	// Print the modified matrix
	for _, row := range matrix {
		fmt.Println(row)
	}
}
