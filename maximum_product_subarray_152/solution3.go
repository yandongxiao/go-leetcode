package maximum_product_subarray_152

func maxProduct(nums []int) int {
	largestValues := make([]int, len(nums))
	smallestValues := make([]int, len(nums))
	largestValues[0] = nums[0]
	smallestValues[0] = nums[0]

	maxValue := largestValues[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			largestValues[i] = max(nums[i], largestValues[i-1]*nums[i])
			smallestValues[i] = min(nums[i], smallestValues[i-1]*nums[i])
		} else {
			largestValues[i] = max(nums[i], smallestValues[i-1]*nums[i])
			smallestValues[i] = min(nums[i], largestValues[i-1]*nums[i])
		}

		if maxValue < largestValues[i] {
			maxValue = largestValues[i]
		}
	}
	return maxValue
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
