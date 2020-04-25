package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	mapping := map[int]int{}
	for i, x := range nums {
		mapping[x] = i
	}

	result := map[[4]int]struct{}{}
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		if x > target && x > 0 {
			break
		}
		for i+4 < len(nums) && nums[i+4] == x {
			i++
		}
		for j := i + 1; j < len(nums); j++ {
			y := nums[j]
			if x+y > target && y > 0 {
				break
			}
			for j+3 < len(nums) && nums[j+3] == y {
				j++
			}
			for k := j + 1; k < len(nums); k++ {
				z := nums[k]
				if x+y+z > target && z > 0 {
					break
				}
				if k+2 < len(nums) && nums[k+2] == z {
					k++
				}
				key := target - (x + y + z)
				if l, ok := mapping[key]; ok && l > k {
					result[[4]int{x, y, z, key}] = struct{}{}
				}
			}
		}
	}

	output := [][]int{}
	for key := range result {
		newKey := key
		output = append(output, newKey[:])
	}
	return output
}
