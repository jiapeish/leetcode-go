package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	low, high := 0, m*n-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if matrix[mid/n][mid%n] == target {
			return true
		} else if matrix[mid/n][mid%n] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
