package _036_valid_sudoku

func isValidSudoku(board [][]byte) bool {
	var seen = make(map[string]bool, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				cell := "(" + string(board[i][j]) + ")"

				// '5' in row 6 is represented as "6(5)"
				if _, ok := seen[string(i)+cell]; ok {
					return false
				} else {
					seen[string(i)+cell] = true
				}

				// '5' in column 6 is represented as "(5)6"
				if _, ok := seen[cell+string(j)]; ok {
					return false
				} else {
					seen[cell+string(j)] = true
				}

				// '5' in middle-left box is represented as "1(5)0"
				if _, ok := seen[string(i/3) + cell + string(j/3)]; ok {
					return false
				} else {
					seen[string(i/3) + cell + string(j/3)] = true
				}
			}
		}
	}

	return true
}
