package main

func isPerfectSquare(num int) bool {
	if num <= 0 {
		return false
	} else if num == 1 {
		return true
	}

	begin := 1
	end := num
	for begin < end {
		mid := begin + (end-begin)/2
		val := mid * mid
		if val == num {
			return true
		} else if val > num {
			end = mid
		} else if val < num {
			begin = mid + 1
		}
	}
	return false
}
