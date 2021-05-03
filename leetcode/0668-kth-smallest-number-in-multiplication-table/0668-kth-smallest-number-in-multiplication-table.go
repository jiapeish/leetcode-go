package _668_kth_smallest_number_in_multiplication_table

import "math"

func findKthNumber(m int, n int, k int) int {
	low, high := 1, m*n
	for low < high {
		mid := low + (high - low) >> 1
		if counterKthNum(m, n, mid) >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func findKthNumber2(m int, n int, k int) int {
	low, high := 1, m*n
	for low <= high {
		mid := low + (high - low) >> 1
		if counterKthNum(m, n, mid) >= k {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func counterKthNum(m int, n int, mid int) int {
	count := 0
	for i := 1; i <= m; i++ {
		count += int(math.Min(float64(mid/i), float64(n)))
	}
	return count
}