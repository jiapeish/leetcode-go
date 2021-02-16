package main

// For example:
// Given the below binary tree and sum = 22,
//
//              5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
//
// return
//
// [
//   [5,4,11,2],
//   [5,8,4,5]
// ]


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var path []int
	var dfs func(root *TreeNode, level int, sum int)

	dfs = func(root *TreeNode, level int, sum int) {
		if root == nil {
			return
		}

		if level < len(path) {
			path[level] = root.Val
		} else {
			path = append(path, root.Val)
		}
		sum -= root.Val

		if root.Left == nil && root.Right == nil && sum == 0 {
			tmp := make([]int, level+1)
			copy(tmp, path)
			res = append(res, tmp)
		}

		dfs(root.Left, level+1, sum)
		dfs(root.Right, level+1, sum)
	}

	dfs(root, 0, targetSum)
	return res
}