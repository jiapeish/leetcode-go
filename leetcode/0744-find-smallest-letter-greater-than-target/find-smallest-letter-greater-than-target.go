package _744_find_smallest_letter_greater_than_target

import "sort"

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	j := sort.Search(n, func(i int) bool {
		return target < letters[i]
	})
	return letters[j%n]
}