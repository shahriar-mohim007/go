package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return true
	}

	// Create adjacency list and indegree array
	adjList := make(map[int][]int)
	indegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		adjList[pre[1]] = append(adjList[pre[1]], pre[0])
		indegree[pre[0]]++
	}

	// Initialize queue with courses having 0 indegree
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// BFS traversal
	c := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		c++

		// Reduce indegree of neighbors
		for _, neighbor := range adjList[current] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Check if all courses can be finished
	return c == numCourses
}

func main() {
	// Test cases
	numCourses1 := 2
	prerequisites1 := [][]int{{1, 0}}

	numCourses2 := 2
	prerequisites2 := [][]int{{1, 0}, {0, 1}}

	fmt.Println("Test Case 1:", canFinish(numCourses1, prerequisites1)) // Expected output: true
	fmt.Println("Test Case 2:", canFinish(numCourses2, prerequisites2)) // Expected output: false
}
