package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Right)
		sum += node.Val
		node.Val = sum
		traversal(node.Left)
	}
	traversal(root)

	return root
}