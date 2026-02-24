package week06

// dp
func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}

	dp := make([]int, len(s))
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			}
			if s[i-1] == ')' {
				if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					if i-dp[i-1]-2 >= 0 {
						dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
