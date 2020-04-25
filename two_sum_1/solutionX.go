package two_sum_1

func twoSum(nums []int, target int) []int {
	mapping := map[int]int
	for i, x := range nums {
		if j, ok := mapping[target-x]; ok {
			return []int{j, i}
		}
		mapping[x] = i
	}
	return nil
}
