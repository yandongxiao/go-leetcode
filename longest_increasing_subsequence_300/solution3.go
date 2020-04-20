package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	lis := make([]int, len(nums))
	for i := range lis {
		lis[i] = 1
	}
	max := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				lis[i] = fmax(lis[i], lis[j]+1)
			}
		}
		if lis[i] > max {
			max = lis[i]
		}
	}
	fmt.Println(lis)

	return max
}

func fmax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}
