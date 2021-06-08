package main

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	dfs(root, nil)
	return root;
}

func dfs(curr *Node, next *Node) {
	if curr == nil {
		return
	}

	curr.Next = next;
	dfs(curr.Left, curr.Right)

	if curr.Next == nil {
		dfs(curr.Right, nil)
	} else {
		dfs(curr.Right, curr.Next.Left)
	}
}

// todo wrong!
func connectIter(root *Node) *Node {
	if root == nil {
		return root
	}

	var pre = root
	var cur *Node

	for pre.Left != nil {
		cur = pre
		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}
		pre = pre.Next
	}

	return root
}