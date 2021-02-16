package main

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return transMid2Tree(head, nil)
}

func transMid2Tree(begin *ListNode, end *ListNode) *TreeNode {
	if begin == end {
		return nil
	}

	if begin.Next == end {
		return &TreeNode{Val: begin.Val}
	}

	fast, slow := begin, begin
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow

	return &TreeNode{
		Val: mid.Val,
		Left: transMid2Tree(begin, mid),
		Right: transMid2Tree(mid.Next, end),
	}

}



