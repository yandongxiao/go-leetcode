package number_of_1_bits_191

func hammingWeight(num uint32) int {
	counter := uint32(0)
	for num > 0 {
		counter += num % 2
		num /= 2
	}
	return int(counter)
}
