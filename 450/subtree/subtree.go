package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func main() {
	// Example Usage
	// Build the binary tree for root and subRoot
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 0}

	subRoot := &TreeNode{Val: 4}
	subRoot.Left = &TreeNode{Val: 1}
	subRoot.Right = &TreeNode{Val: 2}

	// Check if subRoot is a subtree of root
	fmt.Println(isSubtree(root, subRoot)) // Output: true
}
