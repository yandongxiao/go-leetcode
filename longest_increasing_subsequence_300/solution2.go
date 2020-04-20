package longest_increasing_subsequence_300

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	lis := make([]int, len(nums))
	lis[0] = 1
	max := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			lis[i] = lis[i-1] + 1
		} else {
			lis[i] = 1
		}
		if max < lis[i] {
			max = lis[i]
		}
	}
	fmt.Println(lis)
	return max
}
