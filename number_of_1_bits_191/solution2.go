package number_of_1_bits_191

func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		result += int(num % 2)
		num = num / 2
	}
	return result
}
