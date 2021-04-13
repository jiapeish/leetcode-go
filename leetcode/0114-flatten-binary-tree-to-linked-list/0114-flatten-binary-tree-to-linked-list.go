package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	list, node := []int{}, &TreeNode{}
	list = preorder(root)
	node = root
	for i := 1; i < len(list); i++ {
		node.Left = nil
		node.Right = &TreeNode{
			Val:   list[i],
			Left:  nil,
			Right: nil,
		}
		node = node.Right
	}
	return
}

func preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var output []int
	output = append(output, root.Val)
	output = append(output, preorder(root.Left)...)
	output = append(output, preorder(root.Right)...)
	return output
}