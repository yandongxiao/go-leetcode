package coin_change_322

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	result := make([]int, amount+1)
	for i := 0; i < len(coins); i++ {
		v := coins[i]
		if v <= amount {
			result[v] = 1
		}
	}

	for i := 1; i <= amount; i++ {
		if result[i] != 0 {
			continue
		}

		for _, x := range coins {
			if i > x && result[i-x] > 0 {
				if result[i] == 0 {
					result[i] = result[i-x] + 1
				} else {
					result[i] = min(result[i], result[i-x]+1)
				}
			}
		}
	}

	if result[amount] == 0 {
		return -1
	}
	return result[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
