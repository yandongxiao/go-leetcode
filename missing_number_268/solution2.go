package missing_number_268

func missingNumber(nums []int) int {
	n := len(nums)
	expectedSum := n * (n + 1) / 2
	actualSum := 0
	for _, x := range nums {
		actualSum += x
	}
	return expectedSum - actualSum
}
