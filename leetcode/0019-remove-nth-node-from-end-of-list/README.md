package _019_remove_nth_node_from_end_of_list

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	slow, fast := dummy, dummy
	slow.Next = head

	for i := 1; i <= n + 1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
