package week07

func isValidSudoku(board [][]byte) bool {
	if board == nil || len(board) != 9 {
		return false
	}

	for i := 0; i < len(board); i++ {
		if len(board[i]) != 9 {
			return false
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' && !isValid(board, i, j, board[i][j]) {
				return false
			}
		}
	}

	return true
}

func isValid(board [][]byte, row, col int, val byte) bool {
	for i := 0; i < 9; i++ {
		if i != row && board[i][col] != '.' && board[i][col] == val {
			return false
		}
		if i != col && board[row][i] != '.' && board[row][i] == val {
			return false
		}

		x := 3*(row/3) + i/3
		y := 3*(col/3) + i%3
		if x == row && y == col {
			continue
		}
		if board[x][y] != '.' && board[x][y] == val {
			return false
		}
	}

	return true
}
