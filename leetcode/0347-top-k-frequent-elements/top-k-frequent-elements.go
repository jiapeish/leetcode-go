package _347_top_k_frequent_elements

import "sort"

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	freq := make(map[int]int, len(nums))
	count := make([]int, 0, len(nums))

	for _, num := range nums {
		freq[num]++
	}

	for _, f := range freq {
		count = append(count, f)
	}

	sort.Ints(count)
	min := count[len(count) - k]

	for num, f := range freq {
		if f >= min {
			res = append(res, num)
		}
	}

	return res
}