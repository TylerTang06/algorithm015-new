package week09

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	obstacleGrid[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
		} else {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			obstacleGrid[0][j] = 0
		} else {
			obstacleGrid[0][j] = obstacleGrid[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}

	return obstacleGrid[m-1][n-1]
}
