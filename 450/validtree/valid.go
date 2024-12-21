package main

import "fmt"

func ValidTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = false
	}
	queue := make([]int, 0)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if visited[cur] {
			return false
		}
		visited[cur] = true
		for _, w := range graph[cur] {
			if !visited[w] {
				queue = append(queue, w)
			}
		}
	}
	return true
}

func main() {
	// Test cases
	n1 := 5
	edges1 := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}

	n2 := 5
	edges2 := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}

	fmt.Println("Test Case 1:", ValidTree(n1, edges1)) // Expected output: true
	fmt.Println("Test Case 2:", ValidTree(n2, edges2)) // Expected output: false
}
