package best_time_to_buy_and_sell_stock_121

func maxProfit(prices []int) int {
	maxProfit := 0
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-buy)
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
