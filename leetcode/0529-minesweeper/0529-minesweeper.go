package main

var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]

	// if click 'M', just change it to 'X', and return
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}

	// otherwise, click 'E', use bfs
	return bfs(board, x, y)
}

// if we click 'E' at (x,y), need to modify all 'E' to 'B' adj to (x,y), and return the board
func bfs(board [][]byte, x, y int) [][]byte {
	m, n := len(board), len(board[0])
	visited := make([]bool, m*n)
	xqueue := make([]int, 0)
	yqueue := make([]int, 0)

	xqueue = append(xqueue, x)
	yqueue = append(yqueue, y)
	visited[x*n+y] = true

	for len(xqueue) > 0 {
		x = xqueue[0]
		y = yqueue[0]
		xqueue = xqueue[1:]
		yqueue = yqueue[1:]

		bombs := getBombs(board, x, y)

		// with at least one adjacent mine
		if bombs > 0 {
			board[x][y] = '0' + byte(bombs)
			continue
		}

		// with no adjacent mines
		board[x][y] = 'B'

		for k := 0; k < 8; k++ {
			i := x + dx[k]
			j := y + dy[k]

			if 0 <= i && i < m && 0 <= j && j < n &&
				board[i][j] == 'E' && !visited[i*n+j] {
				xqueue = append(xqueue, i)
				yqueue = append(yqueue, j)
				visited[i*n+j] = true
			}
		}
	}

	return board
}

func getBombs(board [][]byte, x, y int) int {
	m, n := len(board), len(board[0])
	bombs := 0

	for k := 0; k < 8; k++ {
		i := x + dx[k]
		j := y + dy[k]
		if 0 <= i && i < m && 0 <= j && j < n &&
			board[i][j] == 'M' {
			bombs++
		}
	}
	return bombs
}
