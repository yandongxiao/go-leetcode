package main

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	begin := 1
	end := x
	var mid, val int
	for begin < end {
		mid = begin + (end-begin)/2
		val = mid * mid
		if val > x {
			end = mid
		} else if val == x || (mid+1)*(mid+1) > x {
			return mid
		} else {
			begin = mid + 1
		}
	}
	return -1
}
