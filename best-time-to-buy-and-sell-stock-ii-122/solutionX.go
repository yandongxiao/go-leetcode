package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	buy := prices[0]
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > buy {
			sum += prices[i] - buy
		}
		buy = prices[i]
	}
	return sum
}
