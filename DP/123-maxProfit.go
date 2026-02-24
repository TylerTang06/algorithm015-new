package week06

import (
	"math"
	"sort"
)

func maxProfitTwoTimes(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	profit := [2][3]int{
		[3]int{0, -prices[0], 0},                            // profit[0][k] 0 表示可买入，k表示第k次操作
		[3]int{math.MinInt32, math.MinInt32, math.MinInt32}, // profit[1][k] 1 表示可卖出
	}
	for i := 0; i < len(prices); i++ {
		// profit[0][0] = profit[0][0] // 一直不买不卖的情况
		if profit[1][0] < profit[0][0]-prices[i] { // 第一次买入
			profit[1][0] = profit[0][0] - prices[i]
		}
		if profit[0][1] < profit[1][0]+prices[i] { // 第一次买入后，max(不动,卖出)
			profit[0][1] = profit[1][0] + prices[i]
		}
		if profit[1][1] < profit[0][1]-prices[i] { // 第一次买入后，max(不动,继续买入)
			profit[1][1] = profit[0][1] - prices[i]
		}
		if profit[0][2] < profit[1][1]+prices[i] { // 第二次买入后，max(不动,卖出)
			profit[0][2] = profit[1][1] + prices[i]
		}
	}

	ints := []int{profit[0][0], profit[0][1], profit[0][2]}
	sort.Ints(ints)
	return ints[2]
}
