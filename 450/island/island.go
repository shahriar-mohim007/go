package main

import "fmt"

func bfs(grid [][]byte, i, j int) {
	q := make([][2]int, 0)
	q = append(q, [2]int{i, j})
	gridRow := len(grid)
	gridColumn := len(grid[i])
	dir := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	grid[i][j] = '0'
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, direction := range dir {
			ni, nj := cur[0]+direction[0], cur[1]+direction[1]
			if ni >= 0 && ni < gridRow && nj >= 0 && nj < gridColumn && grid[ni][nj] == '1' {
				grid[ni][nj] = '0'
				q = append(q, [2]int{ni, nj})
			}
		}
	}
}

func numIslands(grid [][]byte) int {

	island := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				island++
				bfs(grid, i, j)
			}
		}
	}
	return island
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	result := numIslands(grid)
	fmt.Println("Number of islands:", result)
}
