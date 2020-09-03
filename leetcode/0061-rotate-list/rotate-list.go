package _061_rotate_list

/**
* Definition for singly-linked list.
*/

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	tail := head
	for i := 0; i < k; i++ {
		if tail.Next == nil {
			return rotateRight(head, k%(i+1))
		}
		tail = tail.Next
	}

	newTail := head
	for tail.Next != nil {
		tail, newTail = tail.Next, newTail.Next
	}

	newHead := newTail.Next
	newTail.Next, tail.Next = nil, head

	return newHead
}
