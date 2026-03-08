package week09

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
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
			} else { // s[i-1] == ')'
				if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					if i-dp[i-1]-2 >= 0 {
						dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
		}
		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}
