package week09

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// binary search
	dp := make([]int, len(nums))
	cur := 0
	for _, num := range nums {
		l, r := 0, cur
		for l < r {
			mid := l + (r-l)/2
			if dp[mid] < num {
				l = mid + 1
			} else {
				r = mid
			}
		}
		dp[l] = num
		if l == cur {
			cur++
		}
	}

	return cur
}
