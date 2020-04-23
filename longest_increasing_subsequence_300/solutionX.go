package longest_increasing_subsequence_300

func lengthOfLIS(nums []int) int {
	result := make([]int, len(nums))
	for i := range result {
		result[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				result[i] = max(result[i], result[j]+1)
			}
		}
	}

	maxVal := result[0]
	for i := 1; i < len(result); i++ {
		maxVal = max(result[i], maxVal)
	}
	return maxVal
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
