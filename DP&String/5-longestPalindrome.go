package week09

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	// dp[i][j] = dp[i+1][j-1] if s[i] == s[j]
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	start, end := 0, 0

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}

			if dp[i][j] == true && j-i > end-start {
				start, end = i, j
			}
		}
	}

	return s[start : end+1]
}

func longestPalindrome1(s string) string {
	if len(s) <= 1 {
		return s
	}

	var longestStr string
	for i := 0; i < len(s)-1; i++ {
		oddStr := centerSpread(s, i, i)
		eventStr := centerSpread(s, i, i+1)
		if len(longestStr) < len(oddStr) {
			longestStr = oddStr
		}
		if len(longestStr) < len(eventStr) {
			longestStr = eventStr
		}
	}

	return longestStr
}

func centerSpread(s string, left, right int) string {
	if len(s) <= 1 {
		return s
	}

	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}

	return s[left+1 : right]
}
