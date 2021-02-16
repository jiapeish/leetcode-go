package main

// preorder = [3,9,20,15,7]
// inorder = [9,3,15,20,7]

//   3
//  / \
// 9  20
//   /  \
//  15   7

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	res := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	if len(preorder) == 1 {
		return res
	}

	idx := indexOf(res.Val, inorder)

	res.Left = buildTree(preorder[1:1+idx], inorder[:idx])
	res.Right = buildTree(preorder[1+idx:], inorder[idx+1:])
	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return 0
}