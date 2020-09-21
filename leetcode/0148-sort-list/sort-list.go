package _148_sort_list

type ListNode struct {
	Val int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left, right := partition(head)
	return merge(sortList(left), sortList(right))
}

func partition(head *ListNode) (left, rigth *ListNode) {

	slow, fast := head, head
	var slowPrev *ListNode

	for fast != nil && fast.Next != nil {
		slowPrev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	slowPrev.Next = nil
	left, rigth = head, slow
	return
}

func merge(left, right *ListNode) *ListNode {

	node := &ListNode{}
	head := node

	for left != nil && right != nil {
		if left.Val < right.Val {
			node.Next = left
			left = left.Next
		} else {
			node.Next = right
			right = right.Next
		}
		node = node.Next
	}

	if left != nil {
		node.Next = left
	}

	if right != nil {
		node.Next = right
	}

	return head.Next
}