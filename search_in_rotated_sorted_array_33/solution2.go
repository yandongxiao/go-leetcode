package search_in_rotated_sorted_array_33

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	} else if nums[0] <= nums[len(nums)-1] {
		return binarySearch(nums, target)
	}

	mid := findPivot(nums)
	fmt.Println(mid)
	if target >= nums[0] {
		return binarySearch(nums[:mid], target)
	}
	idx := binarySearch(nums[mid:], target)
	if idx == -1 {
		return -1
	}
	return mid + idx
}

func binarySearch(nums []int, target int) int {
	start := 0
	end := len(nums)
	for start < end {
		if start+1 == end {
			if nums[start] == target {
				return start
			}
			return -1
		}
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func findPivot(nums []int) int {
	start := 0
	end := len(nums)

	for start < end {
		mid := start + (end-start)/2
		if mid-1 >= 0 && nums[mid-1] > nums[mid] {
			return mid
		}
		if mid+1 < len(nums) && nums[mid] > nums[mid+1] {
			return mid + 1
		}

		if nums[mid] > nums[start] {
			start = mid + 1
		} else if nums[mid] < nums[end-1] {
			end = mid
		}
	}
	return -1
}
