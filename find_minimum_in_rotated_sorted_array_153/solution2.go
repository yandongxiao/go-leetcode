package find_minimum_in_rotated_sorted_array_153

func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}

	begin := 0
	end := len(nums) - 1
	for begin < end {
		mid := begin + (end-begin)/2
		if mid-1 >= 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if mid+1 < len(nums) && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid] > nums[begin] {
			begin = mid
		} else if nums[mid] < nums[end] {
			end = mid
		}
	}
	return -1
}
