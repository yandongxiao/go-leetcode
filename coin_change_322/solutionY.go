package coin_change_322

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// 第i个元素表示：至少需要result[i]个硬币，它们的和等于i
	result := make([]int, amount+1)
	for i := range result {
		result[i] = -1
	}
	for _, x := range coins {
		if x <= amount {
			result[x] = 1
		}
	}

	for i := 1; i <= amount; i++ {
		if result[i] != -1 {
			continue
		}

		for _, x := range coins {
			if i-x >= 0 && result[i-x] != -1 {
				if result[i] == -1 {
					result[i] = result[i-x] + 1
				} else {
					result[i] = min(result[i], result[i-x]+1)
				}
			}
		}
	}
	return result[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
