package best_time_to_buy_and_sell_stock_121

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	buy := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		maxProfit = max(maxProfit, prices[i]-buy)
		if prices[i] < buy {
			buy = prices[i]
		}
	}
	return maxProfit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
