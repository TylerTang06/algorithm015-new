package week01

/*
42. 接雨水

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例1:
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

提示：

n == height.length
1 <= n <= 2 * 10^4
0 <= height[i] <= 10^5
*/
func trap(height []int) int {
	if height == nil || len(height) <= 2 {
		return 0
	}

	maxL, maxR, sum := 0, 0, 0
	left, right := 1, len(height)-2
	for i := 1; i < len(height)-1; i++ {
		if height[left-1] < height[right+1] {
			if maxL < height[left-1] {
				maxL = height[left-1]
			}
			if maxL > height[left] {
				sum += (maxL - height[left])
			}
			left++
		} else {
			if maxR < height[right+1] {
				maxR = height[right+1]
			}
			if maxR > height[right] {
				sum += (maxR - height[right])
			}
			right--
		}
	}

	return sum
}
