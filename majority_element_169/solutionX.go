package majority_element_169

func majorityElement(nums []int) int {
	element := nums[0]
	counter := 1
	for i := 1; i < len(nums); i++ {
		if counter == 0 {
			element = nums[i]
			counter++
			continue
		}
		if element == nums[i] {
			counter++
		} else {
			counter--
		}
	}
	return element
}
