package maximum_product_subarray_152

func maxProduct(nums []int) int {
	maxProduct := nums[0]
	minCurrent := nums[0]
	maxCurrent := nums[0]

	for i := 1; i < len(nums); i++ {
		x := nums[i]
		if x < 0 {
			maxCurrent, minCurrent = minCurrent, maxCurrent
		}

		maxCurrent = max(x, maxCurrent*x)
		minCurrent = min(x, minCurrent*x)
		if maxProduct < maxCurrent {
			maxProduct = maxCurrent
		}
	}
	return maxProduct
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
