package week06

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || len(obstacleGrid[0]) == 0 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}

			obstacleGrid[i][j] = 1
			if i == 0 && j > 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1]
			}

			if j == 0 && i > 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j]
			}

			if i > 0 && j > 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}

	return obstacleGrid[m-1][n-1]
}
