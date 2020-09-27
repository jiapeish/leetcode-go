package _024_swap_nodes_in_pairs

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// head next(newHead)
	//  |    |
	//  v    v
	//  1 -> 2 -> 3 -> 4
	//  2 -> 1 -> 4 -> 3
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}