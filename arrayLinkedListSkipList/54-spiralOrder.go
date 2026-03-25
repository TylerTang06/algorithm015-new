package week01

/*
54. 螺旋矩阵
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var directs = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	m, n := len(matrix), len(matrix[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	res := make([]int, m*n)
	var x, y, d int
	for i := 0; i < m*n; i++ {
		res[i] = matrix[x][y]
		visited[x][y] = true
		nx, ny := x+directs[d][0], y+directs[d][1]
		if nx < 0 || nx >= m || ny < 0 || ny >= n || visited[nx][ny] {
			d = (d + 1) % 4
		}
		x += directs[d][0]
		y += directs[d][1]
	}
	return res
}
