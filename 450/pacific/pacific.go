package main

import "fmt"

var (
	dir = [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
)

func dfs(grid [][]int, visited map[[2]int]bool, i, j, height int, row, col int) {
	if visited[[2]int{i, j}] || i < 0 || j < 0 || i >= row || j >= col || grid[i][j] < height {
		return
	}
	visited[[2]int{i, j}] = true
	for _, direction := range dir {
		ni, nj := i+direction[0], j+direction[1]
		dfs(grid, visited, ni, nj, grid[i][j], row, col)
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	row, col := len(heights), len(heights[0])
	pac := make(map[[2]int]bool)
	alt := make(map[[2]int]bool)

	// Process top and bottom edges
	for c := 0; c < col; c++ {
		dfs(heights, pac, 0, c, heights[0][c], row, col)
		dfs(heights, alt, row-1, c, heights[row-1][c], row, col)
	}

	// Process left and right edges
	for r := 0; r < row; r++ {
		dfs(heights, pac, r, 0, heights[r][0], row, col)
		dfs(heights, alt, r, col-1, heights[r][col-1], row, col)
	}

	// Compile results
	result := make([][]int, 0)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if pac[[2]int{r, c}] && alt[[2]int{r, c}] {
				result = append(result, []int{r, c})
			}
		}
	}
	return result
}

func main() {
	heights := [][]int{
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6},
		{3, 4, 5, 6, 7},
		{4, 5, 6, 7, 8},
		{5, 6, 7, 8, 9},
	}
	result := pacificAtlantic(heights)
	fmt.Println(result)
}
