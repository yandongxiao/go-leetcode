package maximum_product_subarray_152

// 方法一
func maxProduct(nums []int) int {
	// 乘积中的最后一个元素是nums[i], 且乘积最大
	largetest := make([]int, len(nums))
	// 乘积中的最后一个元素是nums[i], 且乘积最小
	smallest := make([]int, len(nums))
	largetest[0] = nums[0]
	smallest[0] = nums[0]
	maxValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			largetest[i] = max(largetest[i-1]*nums[i], nums[i])
			smallest[i] = min(smallest[i-1]*nums[i], nums[i])
		} else {
			largetest[i] = max(smallest[i-1]*nums[i], nums[i])
			smallest[i] = min(largetest[i-1]*nums[i], nums[i])
		}
		if maxValue < largetest[i] {
			maxValue = largetest[i]
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
