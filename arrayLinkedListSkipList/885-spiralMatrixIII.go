package week01

/*
885. Spiral Matrix III
在 rows x cols 的网格上，你从单元格 (rStart, cStart) 面朝东面开始。网格的西北角位于第一行第一列，网格的东南角位于最后一行最后一列。

你需要以顺时针按螺旋状行走，访问此网格中的每个位置。每当移动到网格的边界之外时，需要继续在网格之外行走（但稍后可能会返回到网格边界）。

最终，我们到过网格的所有 rows x cols 个空间。

按照访问顺序返回表示网格位置的坐标列表。

输入：rows = 1, cols = 4, rStart = 0, cStart = 0
输出：[[0,0],[0,1],[0,2],[0,3]]
*/
func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	// 四个方向：东、南、西、北
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	result := make([][]int, 0, rows*cols)
	x, y := rStart, cStart
	step := 1   // 当前步数
	dirIdx := 0 // 当前方向索引

	for len(result) < rows*cols {
		// 每个步数走两次（东、南是 step，西、北也是 step）
		for i := 0; i < 2; i++ {
			for j := 0; j < step; j++ {
				// 如果在网格内，添加坐标
				if x >= 0 && x < rows && y >= 0 && y < cols {
					result = append(result, []int{x, y})
				}

				// 如果已收集完所有坐标，提前返回
				if len(result) == rows*cols {
					return result
				}

				// 按当前方向移动
				x += dirs[dirIdx][0]
				y += dirs[dirIdx][1]
			}
			// 转向下一个方向
			dirIdx = (dirIdx + 1) % 4
		}
		// 步数递增
		step++
	}

	return result
}
