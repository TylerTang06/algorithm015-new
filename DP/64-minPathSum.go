package week06

func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// dp[i][j] = min{dp[i-1][j], dp[i][j-1]} + grid[i][j]
	// so we can do it by grid[i][j] += min{grid[i-1][j], grid[i][j-1]}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 || j == 0 {
				if j > 0 {
					grid[i][j] += grid[i][j-1]
				}
				if i > 0 {
					grid[i][j] += grid[i-1][j]
				}
			} else {
				if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] += grid[i][j-1]
				} else {
					grid[i][j] += grid[i-1][j]
				}
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
