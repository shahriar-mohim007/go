package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode, result *[]int, k int) {
	if root == nil {
		return
	}
	if k == len(*result) {
		return
	}
	InorderTraversal(root.Left, result, k)
	*result = append(*result, root.Val)
	InorderTraversal(root.Right, result, k)
}
func kthSmallest(root *TreeNode, k int) int {
	result := make([]int, 0)
	InorderTraversal(root, &result, k)
	fmt.Println(result)
	return result[k-1]

}

func main() {
	// Example usage:
	// Create a sample BST
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 2}

	k := 1
	fmt.Println(kthSmallest(root, k)) // Output: 1 (the 1st smallest element)
}
