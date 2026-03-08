package week04

import "container/list"

// BFS
func updateBoard(board [][]byte, click []int) [][]byte {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return board
	}

	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	// init structures
	visted := map[[2]int]bool{}
	visted[[2]int{click[0], click[1]}] = true
	myQue := list.New()
	myQue.PushBack(click)
	dx := []int{-1, 0, 1, 1, 1, 0, -1, -1}
	dy := []int{1, 1, 1, 0, -1, -1, -1, 0}
	// BFS
	for myQue.Len() > 0 {
		p := myQue.Front().Value.([]int)
		myQue.Remove(myQue.Front())

		count := 0
		// count (0-8)
		for i := 0; i < 8; i++ {
			newP := []int{p[0] + dx[i], p[1] + dy[i]}
			if newP[0] < 0 || newP[0] >= len(board) || newP[1] < 0 || newP[1] >= len(board[0]) {
				continue
			}

			if board[newP[0]][newP[1]] == 'M' {
				count++
				continue
			}
		}
		// if p around by 'M'(count > 0)
		if count > 0 {
			board[p[0]][p[1]] = byte('0' + count)
			continue
		}
		board[p[0]][p[1]] = 'B'

		// recursion if count == 0
		for i := 0; i < 8; i++ {
			newP := []int{p[0] + dx[i], p[1] + dy[i]}
			if newP[0] < 0 || newP[0] >= len(board) || newP[1] < 0 || newP[1] >= len(board[0]) {
				continue
			}
			if _, ok := visted[[2]int{newP[0], newP[1]}]; ok {
				continue
			}

			board[newP[0]][newP[1]] = 'B'
			myQue.PushBack(newP)
			visted[[2]int{newP[0], newP[1]}] = true
		}
	}

	return board
}
