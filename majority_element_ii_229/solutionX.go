package majority_element_ii_229

func majorityElement(nums []int) []int {
	val1, val2 := 1, 2
	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if val1 == num {
			cnt1++
		} else if val2 == num {
			cnt2++
		} else if cnt1 == 0 {
			val1 = num
			cnt1++
		} else if cnt2 == 0 {
			val2 = num
			cnt2++
		} else {
			cnt1--
			cnt2--
		}
	}

	cnt1 = 0
	cnt2 = 0
	for _, num := range nums {
		if num == val1 {
			cnt1++
		} else if num == val2 {
			cnt2++
		}
	}

	var result []int
	if cnt1 > len(nums)/3 {
		result = append(result, val1)
	}
	if cnt2 > len(nums)/3 {
		result = append(result, val2)
	}
	return result
}
