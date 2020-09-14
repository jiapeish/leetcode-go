package _274_h_index

import "sort"

func hIndex(citations []int) int {
	num := len(citations)
	bucket := make([]int, num+1)

	for _, c := range citations {
		if c >= num {
			bucket[num]++
		} else {
			bucket[c]++
		}
	}

	count := 0
	for i := num; i >= 0; i-- {
		count += bucket[i]
		if count >= i {
			return i
		}
	}
	return 0
}

func hIndex2(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	num := len(citations)
	lo, hi := 0, num-1
	var mid int

	for lo <= hi {
		mid = lo + (hi - lo) / 2
		if citations[mid] > mid {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}