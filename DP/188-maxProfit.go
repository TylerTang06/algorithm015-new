package week06

import "sort"

func maxProfitKthTimes(k int, prices []int) int {
	if prices == nil || len(prices) <= 1 || k <= 0 {
		return 0
	}

	// 当k大于等于数组长度一半时, 问题退化为贪心问题此时采用贪心方法解决可以大幅提升时间性能限
	if k >= len(prices)/2 {
		return maxProfitGreedy(prices)
	}

	profit := make([][][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		for kk := 0; kk <= k; kk++ {
			// profit[i][kk][0] 第kk次卖出后的状态；profit[i][kk][1] 第kk次买入后的状态
			profit[i] = append(profit[i], [2]int{0, -prices[0]})
		}
	}

	res := 0
	for i := 1; i < len(prices); i++ {
		ints := []int{}
		for kk := 0; kk <= k; kk++ {
			// 第i天 已完成kk次交易, 卖出一次算一次交易
			if kk != 0 && profit[i-1][kk][0] < profit[i-1][kk-1][1]+prices[i] {
				profit[i][kk][0] = profit[i-1][kk-1][1] + prices[i]
			} else {
				profit[i][kk][0] = profit[i-1][kk][0]
			}

			// 第i天 已完成kk次交易，仍可买入的状态
			if profit[i-1][kk][1] < profit[i-1][kk][0]-prices[i] {
				profit[i][kk][1] = profit[i-1][kk][0] - prices[i]
			} else {
				profit[i][kk][1] = profit[i-1][kk][1]
			}

			ints = append(ints, profit[i][kk][0])
		}
		sort.Ints(ints)

		if res < ints[k] {
			res = ints[k]
		}
	}

	return res
}

func maxProfitGreedy(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	var profit int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[1+i] {
			profit += prices[1+i] - prices[i]
		}
	}
	return profit
}
