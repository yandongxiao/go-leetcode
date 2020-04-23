package main

func isPowerOfTwo(n int) bool {
	counter := 0
	for n > 0 {
		counter += n % 2
		n /= 2
	}
	return counter == 1
}
