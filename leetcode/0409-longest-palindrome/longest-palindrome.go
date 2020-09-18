package _409_longest_palindrome

func longestPalindrome(s string) int {
	var record [256]int
	for i := range s {
		record[s[i]]++
	}

	var length int
	var hasOdd int
	for _, count := range record {
		if count == 0 {
			continue
		}

		if count & 1 == 0 {
			length += count
		} else {
			hasOdd = 1
			length += count - 1
		}
	}
	return length + hasOdd
}