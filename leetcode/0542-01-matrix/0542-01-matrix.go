package main

import (
	"math"
)

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

// DFS
func updateMatrixDFS(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// cell is 1, and has no 0 around itself
			if matrix[i][j] == 1 && adjZero(matrix, i, j) == false {
				matrix[i][j] = math.MaxInt32
			}

		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				dfsMatrix(matrix, i, j, -1)
			}
		}
	}

	return matrix
}

func adjZero(matrix [][]int, row, col int) bool {
	if row > 0 && matrix[row-1][col] == 0 {
		return true
	}

	if col > 0 && matrix[row][col-1] == 0 {
		return true
	}

	if row+1 < len(matrix) && matrix[row+1][col] == 0 {
		return true
	}

	if col+1 < len(matrix[0]) && matrix[row][col+1] == 0 {
		return true
	}

	return false
}

func dfsMatrix(matrix [][]int, row, col, val int) {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) ||
		matrix[row][col] <= val {
		return
	}

	if val > 0 {
		matrix[row][col] = val
	}

	dfsMatrix(matrix, row-1, col, matrix[row][col]+1)
	dfsMatrix(matrix, row, col-1, matrix[row][col]+1)
	dfsMatrix(matrix, row+1, col, matrix[row][col]+1)
	dfsMatrix(matrix, row, col+1, matrix[row][col]+1)
}
