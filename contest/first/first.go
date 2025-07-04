package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())
	for t := 1; t <= T; t++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		graph := make(map[int][]int)
		for i := 0; i < n; i++ {
			scanner.Scan()
			parts := strings.Fields(scanner.Text())
			u, _ := strconv.Atoi(parts[0])
			v, _ := strconv.Atoi(parts[1])
			graph[u] = append(graph[u], v)
			graph[v] = append(graph[v], u)
		}
		visited := make(map[int]bool)
		maxteamsize := 0

		for player := range graph {
			if !visited[player] {
				teamsize := bfscomponentsize(graph, visited, player)
				maxteamsize += teamsize
			}
		}
		fmt.Fprintf(writer, "Tournament %d: %d\n", t, maxteamsize)
	}
}

func bfscomponentsize(graph map[int][]int, visited map[int]bool, start int) int {
	queue := []int{start}
	visited[start] = true
	color := make(map[int]int)
	color[start] = 0
	count := [2]int{1, 0}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				color[neighbor] = 1 - color[node]
				count[color[neighbor]]++
				queue = append(queue, neighbor)
			} else if color[neighbor] == color[node] {
				continue
			}
		}
	}
	return max(count[0], count[1])
}
