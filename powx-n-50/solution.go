package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	num := n
	if num < 0 {
		num = -num
	}

	result := myPow(x*x, num/2)
	if num%2 == 1 {
		result *= x
	}

	if n < 0 {
		return 1 / result
	}
	return result
}
