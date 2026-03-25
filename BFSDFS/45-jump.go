package week04

/*
45.跳跃游戏 II
给定一个长度为 n 的 0 索引整数数组 nums。初始位置在下标 0。

每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在索引 i 处，你可以跳转到任意 (i + j) 处：

0 <= j <= nums[i] 且
i + j < n
返回到达 n - 1 的最小跳跃次数。测试用例保证可以到达 n - 1。

示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例 2:

输入: nums = [2,3,0,1,4]
输出: 2
*/

// 4 ms 99.85% O(n)
// 4 MB 100.00% O(1)
func jump(nums []int) int {
	len := len(nums)
	if nums == nil || len == 0 {
		return 0
	}

	endP, maxP, step := 0, 0, 0
	for i := 0; i < len-1; i++ {
		if maxP < nums[i]+i {
			maxP = nums[i] + i
		}
		if i == endP {
			step++
			endP = maxP
		}
	}

	return step
}

/* O(n*n) O(n) 396 ms 4.5 MB
func jump(nums []int) int {
	len := len(nums)
	if nums == nil || len == 0 {
		return -1
	}

	// dp[i] = min(dp[0]+1, dp[1]+1,...,dp[i-1]+1)
	dp := make([]int, len)
	dp[0] = 0
	for i := 1; i < len; i++ {
		dp[i] = len
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				if dp[i] > dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	if dp[len-1] == len {
		return -1
	}

	return dp[len-1]
}
*/
