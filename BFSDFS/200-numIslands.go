package week04

/*
200.岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

示例 1：

输入：grid = [

	['1','1','1','1','0'],
	['1','1','0','1','0'],
	['1','1','0','0','0'],
	['0','0','0','0','0']

]
输出：1
示例 2：

输入：grid = [

	['1','1','0','0','0'],
	['1','1','0','0','0'],
	['0','0','1','0','0'],
	['0','0','0','1','1']

]
输出：3
*/
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			num := 0
			grid, num = isIsland(grid, i, j)
			res += num
		}
	}

	return res
}

func isIsland(grid [][]byte, x, y int) ([][]byte, int) {
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	if grid[x][y] == '0' {
		return grid, 0
	}
	grid[x][y] = '0'
	for i := 0; i < 4; i++ {
		if isVaild(grid, x+dx[i], y+dy[i]) {
			grid, _ = isIsland(grid, x+dx[i], y+dy[i])
		}
	}
	return grid, 1
}

func isVaild(grid [][]byte, x, y int) bool {
	return !(x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0')
}
