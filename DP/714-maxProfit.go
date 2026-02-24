package week06

func maxProfitWithFee(prices []int, fee int) int {
	if prices == nil || len(prices) <= 1 {
		return 0
	}

	profit := make([][2]int, len(prices))
	profit[0][0] = 0
	profit[0][1] = -prices[0]

	res := 0
	for i := 1; i < len(prices); i++ {
		if profit[i-1][0] < profit[i-1][1]+prices[i]-fee {
			profit[i][0] = profit[i-1][1] + prices[i] - fee
		} else {
			profit[i][0] = profit[i-1][0]
		}

		if profit[i-1][1] < profit[i-1][0]-prices[i] {
			profit[i][1] = profit[i-1][0] - prices[i]
		} else {
			profit[i][1] = profit[i-1][1]
		}

		if res < profit[i][0] {
			res = profit[i][0]
		}
	}

	return res
}
