package maximum_subarray_53

import "math"

func maxSubArray(nums []int) int {
	maxSum := math.MinInt32
	maxCurrent := math.MinInt32

	for _, x := range nums {
		if x > maxCurrent+x {
			maxCurrent = x
		} else {
			maxCurrent += x
		}

		if maxSum < maxCurrent {
			maxSum = maxCurrent
		}
	}
	return maxSum
}
