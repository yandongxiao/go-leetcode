package product_of_array_except_self_238

func productExceptSelf(nums []int) []int {
	product := 1
	zero := 0
	for _, x := range nums {
		if x == 0 {
			zero++
		} else {
			product *= x
		}
	}

	for i := range nums {
		if zero > 1 {
			nums[i] = 0
		} else if zero == 1 {
			if nums[i] == 0 {
				nums[i] = product
			} else {
				nums[i] = 0
			}
		} else {
			nums[i] = product / nums[i]
		}
	}

	return nums
}
