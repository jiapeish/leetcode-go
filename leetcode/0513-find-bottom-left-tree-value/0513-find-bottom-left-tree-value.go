package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findBottomLeftValue1(root *TreeNode) int {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]

		if root.Right != nil {
			queue = append(queue, root.Right)
		}

		if root.Left != nil {
			queue = append(queue, root.Left)
		}
	}
	return root.Val
}

// https://leetcode.com/problems/find-bottom-left-tree-value/discuss/98802/Simple-Java-Solution-beats-100.0!
func findBottomLeftValue(root *TreeNode) int {
	return doFindBottomLeftValue(root, 1, []int{0, 0})
}

func doFindBottomLeftValue(root *TreeNode, depth int, res []int) int {
	if res[1] < depth {
		res[0] = root.Val
		res[1] = depth // res[1] represents maxDepth
	}
	if root.Left != nil {
		doFindBottomLeftValue(root.Left, depth+1, res)
	}
	if root.Right != nil {
		doFindBottomLeftValue(root.Right, depth+1, res)
	}
	return res[0]
}