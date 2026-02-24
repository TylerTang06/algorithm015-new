package week04

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
