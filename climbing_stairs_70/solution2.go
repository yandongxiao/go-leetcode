package climbing_stairs_70

func climbStairs(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 3
	}

	return climbStairs(n/2)*climbStairs(n-n/2) + climbStairs(n/2-1)*climbStairs(n-n/2-1)
}
