package week01

/*
59. 螺旋矩阵 II
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
*/

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}

	var directs = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	val := 1
	x, y, d := 0, 0, 0
	for i := 0; i < n*n; i++ {
		res[x][y] = val
		val++
		nx, ny := x+directs[d][0], y+directs[d][1]
		if nx < 0 || nx >= n || ny < 0 || ny >= n || res[nx][ny] != 0 {
			d = (d + 1) % 4
		}
		x += directs[d][0]
		y += directs[d][1]
	}

	return res
}
