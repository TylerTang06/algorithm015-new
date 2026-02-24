package week06

func maxProfitOneTime(prices []int) int {
	if prices == nil || len(prices) <= 1 {
		return 0
	}

	min, res := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if res < prices[i]-min {
			res = prices[i] - min
		}
	}

	return res
}

// func maxProfit(prices []int) int {
// 	if len(prices) == 0 {
// 		return 0
// 	}
//
// 	result := 0
//  // profit[0]不动；profit[1]买入；profit[2]卖出
// 	profit := []int{0, -prices[0], 0}
// 	for i := 1; i < len(prices); i++ {
// 		profit[2] = profit[1] + prices[i]
// 		if profit[1] < profit[0]-prices[i] {
// 			profit[1] = profit[0] - prices[i]
// 		}
// 		ints := []int{result, profit[0], profit[1], profit[2]}
// 		sort.Ints(ints)
// 		result = ints[3]
// 	}
//
// 	return result
// }
