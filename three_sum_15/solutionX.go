package three_sum_15

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	result := map[[3]int]struct{}{}

	mapping := make(map[int]int, len(nums))
	for i, x := range nums {
		mapping[x] = i
	}

	if nums[len(nums)-1] < 0 {
		return nil
	}

	for i := 0; i < len(nums); i++ {
		x := nums[i]
		if x > 0 {
			break
		}
		for i+3 < len(nums) && nums[i+3] == x {
			i++
		}

		for j := i + 1; j < len(nums); j++ {
			y := nums[j]
			if x+y > 0 {
				break
			} else if x+y+nums[len(nums)-1] < 0 {
				continue
			}

			if k, ok := mapping[-(x + y)]; ok && k > j {
				vals := [3]int{x, y, -(x + y)}
				sort.Ints(vals[:])
				result[vals] = struct{}{}
			}
		}
	}

	var output [][]int
	for key := range result {
		newKey := key
		output = append(output, newKey[:])
	}
	return output
}
