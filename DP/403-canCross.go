package week06

import "math"

// 240 ms 23.66%
// 6.6 MB 90.24%
func canCross(stones []int) bool {
	if stones == nil || len(stones) <= 1 {
		return true
	}

	// dp[i] = []int{k1, k2} the ith index of stones, can jump {k1, k1-1,k1+1} and {k2, k2-1, k2+1}
	dp := make([][]int, len(stones))
	dp[0] = []int{1}

	for i := 1; i < len(stones); i++ {
		dp[i] = []int{}
		for j := 0; j < i; j++ {
			for _, k := range dp[j] {
				diff := stones[i] - stones[j]
				if j == 0 {
					if diff == k {
						dp[i] = append(dp[i], diff)
					}
					continue
				}
				if math.Abs(float64(diff-k)) <= 1 {
					dp[i] = append(dp[i], diff)
					break
				}
			}
		}
	}

	// if len(dp[len(stones)-1]) > 0 {
	// 	return true
	// }
	// return false
	return len(dp[len(stones)-1]) > 0
}
