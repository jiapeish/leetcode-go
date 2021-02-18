package main

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 1, 1024)
	queue[0] = root
	remain := 1
	max := math.MinInt32

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		remain--

		if max < node.Val {
			max = node.Val
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		if remain == 0 {
			remain = len(queue)
			res = append(res, max)
			max = math.MinInt32
		}
	}
	return res
}