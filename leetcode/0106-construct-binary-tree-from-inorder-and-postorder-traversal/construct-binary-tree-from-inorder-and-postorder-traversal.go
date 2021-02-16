package main

// preorder = [3,9,20,15,7]
// postorder = [9,15,7,20,3]

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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	res := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  nil,
		Right: nil,
	}

	if len(inorder) == 1 {
		return res
	}

	idx := indexOf(res.Val, inorder)

	res.Left = buildTree(inorder[:idx], postorder[:idx])
	res.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
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

