package _094_binary_tree_inorder_traversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	output := inorderTraversal(root.Left)
	output = append(output, root.Val)
	output = append(output, inorderTraversal(root.Right)...)

	return output
}
