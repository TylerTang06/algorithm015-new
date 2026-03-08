package week04

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
