package _141_linked_list_cycle

type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode)  {
	if head == nil {
		return
	}

	cur := head
	size := 0
	for cur != nil {
		cur = cur.Next
		size++
	}

	cur = head
	for i := 0; i < (size-1) / 2; i++ {
		cur = cur.Next
	}

	//               next     tmp
	//                |        |
	//                v        v
	// head->1->2->3->4    ->  5  ->  6
	//             ^
	//             |
	//             cur
	next := cur.Next
	for next != nil {
		tmp := next.Next
		next.Next = cur
		cur = next
		next = tmp
	}
	end := cur

	//     hNext                     eNext
	//       |                         |
	//       v                         v
	// head->1->2->3<->4    <-  5  <-  6 <- end
	for head != end {
		hNext := head.Next
		eNext := end.Next
		head.Next = end
		end.Next = hNext
		head = hNext
		end = eNext
	}

	end.Next = nil
}