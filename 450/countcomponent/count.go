package main

import "fmt"

var (
	visited = map[int]bool{}
	graph   [][]int
)

func bfs(i int) {
	queue := []int{i}
	visited[i] = true
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[curr] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func countComponents(n int, edges [][]int) int {
	// Initialize graph adjacency list
	graph = make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	// Initialize visited map
	for i := 0; i < n; i++ {
		visited[i] = false
	}

	// Count connected components
	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			bfs(i)
			count++
		}
	}
	return count
}

func main() {
	// Test cases
	n1 := 5
	edges1 := [][]int{{0, 1}, {1, 2}, {3, 4}}
	fmt.Println("Test Case 1:", countComponents(n1, edges1)) // Expected output: 2

	n2 := 4
	edges2 := [][]int{{0, 1}, {2, 3}}
	fmt.Println("Test Case 2:", countComponents(n2, edges2)) // Expected output: 2

	n3 := 6
	edges3 := [][]int{{0, 1}, {1, 2}, {2, 3}, {4, 5}}
	fmt.Println("Test Case 3:", countComponents(n3, edges3)) // Expected output: 2
}
