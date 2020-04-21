package coin_change_322

func coinChange(coins []int, amount int) int {
	result := make(map[int]int, amount)
	for _, x := range coins {
		result[x] = 1
	}
	return helper(coins, amount, result)
}

func helper(coins []int, amount int, result map[int]int) int {
	if amount < 0 {
		return -1
	} else if amount == 0 {
		return 0
	}

	if v, ok := result[amount]; ok {
		return v
	}

	mins := make([]int, len(coins))
	for i, x := range coins {
		mins[i] = helper(coins, amount-x, result)
	}
	v := findMinest(mins)
	if v == -1 {
		result[amount] = -1
	} else {
		result[amount] = v + 1
	}

	return result[amount]
}

func findMinest(mins []int) int {
	minVal := -1
	for i := range mins {
		if mins[i] == -1 {
			continue
		}
		if minVal == -1 {
			minVal = mins[i]
		} else if minVal > mins[i] {
			minVal = mins[i]
		}
	}
	return minVal
}
