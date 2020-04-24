package maximum_subarray_53

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSums := make([]int, len(nums))
	maxSums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if maxSums[i-1] < 0 {
			maxSums[i] = nums[i]
		} else {
			maxSums[i] = maxSums[i-1] + nums[i]
		}
	}

	max := maxSums[0]
	for _, x := range maxSums {
		if max < x {
			max = x
		}
	}
	return max
}
