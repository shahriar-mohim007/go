package main

import "fmt"

func dfs(i, j, index, m, n int, visited map[[2]int]bool, words string, board [][]byte) bool {
	// If we've matched the entire word, return true
	if index == len(words) {
		return true
	}

	// Boundary conditions and character matching check
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != words[index] || visited[[2]int{i, j}] {
		return false
	}

	// Mark the current cell as visited
	visited[[2]int{i, j}] = true

	// Explore all four directions (down, up, right, left)
	if dfs(i+1, j, index+1, m, n, visited, words, board) ||
		dfs(i-1, j, index+1, m, n, visited, words, board) ||
		dfs(i, j+1, index+1, m, n, visited, words, board) ||
		dfs(i, j-1, index+1, m, n, visited, words, board) {
		return true
	}

	// Backtrack: Unmark the current cell as visited
	visited[[2]int{i, j}] = false
	return false
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	// Iterate over each cell in the board to start the DFS
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited := make(map[[2]int]bool)
			// Start DFS from the current cell
			if dfs(i, j, 0, m, n, visited, word, board) {
				return true
			}
		}
	}

	// If no path was found, return false
	return false
}

func main() {
	// Example usage
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word)) // Output: true
}
