package week02

/*
264.丑数2
给你一个整数 n ，请你找出并返回第 n 个 丑数 。

丑数 就是质因子只包含 2、3 和 5 的正整数。



示例 1：

输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
示例 2：

输入：n = 1
输出：1
解释：1 通常被视为丑数。


提示：

1 <= n <= 1690
*/

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	a, b, c := 0, 0, 0

	for i := 1; i < n; i++ {
		dp[i] = min(dp[a]*2, dp[b]*3, dp[c]*5)

		// should remove the case of dp[2] * 3 == dp[3] * 2
		if dp[i] == dp[a]*2 {
			a++
		}
		if dp[i] == dp[b]*3 {
			b++
		}
		if dp[i] == dp[c]*5 {
			c++
		}
	}

	return dp[n-1]
}

func min(x, y, z int) int {
	if x > y {
		x = y
	}
	if x > z {
		return z
	}

	return x
}
