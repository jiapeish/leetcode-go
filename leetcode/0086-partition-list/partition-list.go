package _086_partition_list

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {

	if head == nil || head.Next == nil{
		return head
	}

	lessHead := &ListNode{}
	geHead := &ListNode{
		Val:  0,
		Next: nil,
	}

	lessTail := lessHead
	geTail := geHead

	for head != nil {
		if head.Val < x {
			lessTail.Next = head
			lessTail = lessTail.Next
		} else {
			geTail.Next = head
			geTail = geTail.Next
		}
		head = head.Next
	}

	lessTail.Next = geHead.Next
	geTail.Next = nil

	return lessHead.Next
}