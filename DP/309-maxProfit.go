package week06

import (
	"math"
	"sort"
)

func maxProfitCoolDown(prices []int) int {
	if prices == nil || len(prices) <= 1 {
		return 0
	}

	profit := make([][2][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		// profit[i][0][0] 第i天cooldown；profit[i][1][1] 第i天可卖状态(买入后的状态)
		// profit[i][1][0] 第i天可买入；
		profit[i][0] = [2]int{0, math.MinInt32}
		profit[i][1] = [2]int{0, -prices[0]}
	}

	res := 0
	for i := 1; i < len(prices); i++ {
		profit[i][0][0] = profit[i-1][1][1] + prices[i]
		if profit[i-1][1][1] < profit[i-1][1][0]-prices[i] {
			profit[i][1][1] = profit[i-1][1][0] - prices[i]
		} else {
			profit[i][1][1] = profit[i-1][1][1]
		}
		if profit[i-1][1][0] < profit[i-1][0][0] {
			profit[i][1][0] = profit[i-1][0][0]
		} else {
			profit[i][1][0] = profit[i-1][1][0]
		}

		ints := []int{profit[i][1][0], profit[i][0][0], res}
		sort.Ints(ints)
		res = ints[2]
	}

	return res
}
