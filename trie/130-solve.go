package week07

// DFS, it is not good idea
func solve(board [][]byte) {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return
	}

	visited := map[[2]int]bool{}

	for i := 0; i < len(board); i++ {
		board, visited = solveRec(board, visited, i, 0, 'O')
		board, visited = solveRec(board, visited, i, len(board[0])-1, 'O')
	}

	for j := 1; j < len(board[0])-1; j++ {
		board, visited = solveRec(board, visited, 0, j, 'O')
		board, visited = solveRec(board, visited, len(board)-1, j, 'O')
	}

	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[0])-1; j++ {
			board, visited = solveRec(board, visited, i, j, 'X')
		}
	}
}

func solveRec(board [][]byte, visited map[[2]int]bool, x, y int, replace byte) ([][]byte, map[[2]int]bool) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] == 'X' {
		return board, visited
	}

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}

	if _, ok := visited[[2]int{x, y}]; ok {
		return board, visited
	}

	visited[[2]int{x, y}] = true
	board[x][y] = replace
	for i := 0; i < 4; i++ {
		board, visited = solveRec(board, visited, x+dx[i], y+dy[i], replace)
	}

	return board, visited
}
