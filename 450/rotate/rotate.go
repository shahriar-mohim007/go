package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)

	// Step 1: Transpose the matrix (swap matrix[i][j] with matrix[j][i])
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ { // Start at j = i+1 to avoid redundant swaps
			fmt.Println(i, j, j, i)
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			fmt.Println(matrix)
		}
	}

	// Step 2: Reverse each row
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]

		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(matrix)

	// Print the rotated matrix
	for _, row := range matrix {
		fmt.Println(row)
	}
}
