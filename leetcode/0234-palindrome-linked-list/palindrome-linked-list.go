package _234_palindrome_linked_list

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	values := make([]int, 0, 128)

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	l, r := 0, len(values) - 1
	for l < r {
		if values[l] != values[r] {
			return false
		}
		l++
		r--
	}
	return true
}
