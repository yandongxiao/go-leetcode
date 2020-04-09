package best_time_to_buy_and_sell_stock_121

import "math"

func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		buy := prices[i]
		for j := i + 1; j < len(prices); j++ {
			sell := prices[j]
			if sell-buy > max {
				max = sell - buy
			}
		}
	}
	return max
}

func maxProfit2(prices []int) int {
	buyPrice := math.MaxInt32
	maxProfit := 0
	for _, x := range prices {
		if buyPrice > x {
			buyPrice = x
		} else if maxProfit < x-buyPrice {
			maxProfit = x - buyPrice
		}
	}
	return maxProfit
}
