package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0

	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, num int) {
		if root == nil {
			return
		}

		num = num * 10 + root.Val
		if root.Left == nil && root.Right == nil {
			sum += num
			return
		}

		dfs(root.Left, num)
		dfs(root.Right, num)
	}

	dfs(root, 0)
	return sum
}

