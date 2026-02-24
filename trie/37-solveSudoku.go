package week07

import "fmt"

func solveSudoku(board [][]byte) {
	if len(board) == 0 || board == nil {
		return
	}

	board, ok := solveSudokuRec(board)
	if ok {
		fmt.Println(board)
	}

	return
}

func solveSudokuRec(board [][]byte) ([][]byte, bool) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == '.' {
				for v := byte('1'); v <= byte('9'); v++ {
					if isValidSudokuOne(board, i, j, v) {
						board[i][j] = v
						if board, ok := solveSudokuRec(board); ok {
							return board, true
						}
						board[i][j] = '.'
					}
				}
				return nil, false
			}
		}
	}

	return board, true
}

func isValidSudokuOne(board [][]byte, row, col int, val byte) bool {
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
