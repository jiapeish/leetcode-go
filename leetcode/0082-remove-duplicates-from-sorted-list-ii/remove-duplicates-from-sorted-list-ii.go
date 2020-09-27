package _082_remove_duplicates_from_sorted_list_ii

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// if len <= 1, return
	if head == nil || head.Next == nil {
		return head
	}

	// if head duplicates, remove head
	if head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}

		return deleteDuplicates(head.Next)
	}

	// else, process head.Next recursively
	head.Next = deleteDuplicates(head.Next)
	return head
}
