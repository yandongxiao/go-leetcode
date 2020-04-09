package two_sum_1

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for i, x := range nums {
		left := target - x
		if j, ok := mapping[left]; ok {
			return []int{j, i}
		}
		mapping[x] = i
	}
	return []int{}
}
