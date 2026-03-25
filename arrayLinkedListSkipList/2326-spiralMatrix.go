package week01

import "github.com/TylerTang06/-algorithm015/util"

/*
2326. Spiral Matrix
给你两个整数：m 和 n ，表示矩阵的维数。

另给你一个整数链表的头节点 head 。

请你生成一个大小为 m x n 的螺旋矩阵，矩阵包含链表中的所有整数。链表中的整数从矩阵 左上角 开始、顺时针 按 螺旋 顺序填充。如果还存在剩余的空格，则用 -1 填充。

返回生成的矩阵。

输入：m = 3, n = 5, head = [3,0,2,6,8,1,7,9,4,2,5,5,0]
输出：[[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]
解释：上图展示了链表中的整数在矩阵中是如何排布的。
注意，矩阵中剩下的空格用 -1 填充。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *util.ListNode) [][]int {
	if head == nil {
		return nil
	}

	var directs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = -1
		}
	}

	x, y, d := 0, 0, 0
	for head != nil {
		matrix[x][y] = head.Val
		head = head.Next

		// 计算下一个位置
		xn, yn := x+directs[d][0], y+directs[d][1]

		// 如果下一个位置超出边界或已访问过，则转向
		if xn < 0 || xn >= m || yn < 0 || yn >= n || matrix[xn][yn] != -1 {
			d = (d + 1) % 4
		}

		// 移动到下一个位置
		x += directs[d][0]
		y += directs[d][1]
	}
	return matrix
}
