package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// Create root node from the first element of preorder
	root := &TreeNode{Val: preorder[0]}

	// Find the root index in the inorder traversal
	var rootIndex int
	for i, v := range inorder {
		if v == root.Val {
			rootIndex = i
			break
		}
	}

	// Recursively build the left and right subtrees
	// Left subtree takes elements before rootIndex in inorder
	// Right subtree takes elements after rootIndex in inorder
	root.Left = buildTree(preorder[1:1+rootIndex], inorder[:rootIndex])
	root.Right = buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:])

	return root
}
