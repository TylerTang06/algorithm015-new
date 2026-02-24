package week06

// dp
func maxCoins(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	// dp[i][j] = dp[i][k] + dp[k][j] + nums[i]*nums[j]*nums[k], i < k < j
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}

	for i := len(nums); i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			for k := i + 1; k < j; k++ {
				if dp[i][j] < dp[i][k]+dp[k][j]+nums[i]*nums[j]*nums[k] {
					dp[i][j] = dp[i][k] + dp[k][j] + nums[i]*nums[j]*nums[k]
				}
			}
		}
	}

	return dp[0][len(nums)-1]
}
