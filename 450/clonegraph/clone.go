package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	if _, ok := visited[node]; ok {
		return visited[node]
	}
	clone := &Node{Val: node.Val}
	visited[node] = clone
	for _, n := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfs(n, visited))
	}
	return clone
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	return dfs(node, visited)
}

func main() {
	// Creating a simple graph as an example
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	// Cloning the graph
	clonedGraph := cloneGraph(node1)
	fmt.Println(clonedGraph)
}
