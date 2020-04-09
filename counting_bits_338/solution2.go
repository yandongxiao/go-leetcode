package counting_bits_338

func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		result[i] = (result[i/2]) + (i % 2)
	}
	return result
}
