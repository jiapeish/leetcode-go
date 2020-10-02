package _290_convert_binary_number_in_a_linked_list_to_integer

type ListNode struct {
	Val int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	ans := 0
	for head != nil {
		ans = ans << 1 | head.Val
		head = head.Next
	}
	return ans
}
