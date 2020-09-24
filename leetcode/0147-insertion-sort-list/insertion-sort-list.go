package _141_linked_list_cycle

type ListNode struct {
	Val int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {

	//  prev   next
	// hPrev   head
	//  |	    |
	//  v	    v
	//  L  ->	4    ->    2    ->    1    ->    3		->		5	->		11	->	7
	//          ^          ^
	//          |          |
	//         cur         p

	hPrev := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := head

	for cur != nil && cur.Next != nil {
		p := cur.Next
		if cur.Val <= p.Val {
			cur = p
			continue
		}

		// found p > val, then delete the p after cur
		cur.Next = p.Next
		// need to insert p between "suitable" prev and next
		prev, next := hPrev, hPrev.Next
		// "suitable" means prev.Val < p.Val <= next.Val
		for next.Val < p.Val {
			prev = next
			next = next.Next
		}
		// insert p
		prev.Next = p
		p.Next = next
	}

	return hPrev.Next
}