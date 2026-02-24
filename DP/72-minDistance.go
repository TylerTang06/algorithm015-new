package week06

import (
	"sort"
)

func minDistance(word1 string, word2 string) int {
	if word1 == "" {
		return len(word2)
	}
	if word2 == "" {
		return len(word1)
	}

	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				ints := []int{dp[i-1][j], dp[i][j-1], dp[i-1][j-1]}
				sort.Ints(ints)
				dp[i][j] = ints[0] + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}
