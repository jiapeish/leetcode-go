package main

import "math"

func updateMatrixDP(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	max := m * n

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}

			matrix[i][j] = max

			// update from up and left
			if i - 1 >= 0 {
				matrix[i][j] = min(matrix[i][j], matrix[i-1][j]+1)
			}

			if j - 1 >= 0 {
				matrix[i][j] = min(matrix[i][j], matrix[i][j-1]+1)
			}

		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
			}

			// update from down and right
			if i + 1 < m {
				matrix[i][j] = min(matrix[i][j], matrix[i+1][j]+1)
			}

			if j + 1 < n {
				matrix[i][j] = min(matrix[i][j], matrix[i][j+1]+1)
			}
		}
	}
	return matrix
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// BFS
func updateMatrixBFS(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	queue := make([][]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				matrix[i][j] = math.MaxInt32
			}
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			r := cell[0] + dir[0]
			c := cell[1] + dir[1]
			if r < 0 || r >= m || c < 0 || c >= n ||
				matrix[r][c] <= matrix[cell[0]][cell[1]] + 1 {
				continue
			}

			matrix[r][c] = matrix[cell[0]][cell[1]] + 1
			queue = append(queue, []int{r, c})

		}
	}

	return matrix
}











