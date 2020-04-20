package climbing_stairs_70

// 方法一
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

// 方法二
func climbStairs(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	return climbStairs(n/2)*climbStairs(n-n/2) + climbStairs(n/2-1)*climbStairs(n-n/2-1)
}

// 方法三
func climbStairs(n int) int {
	result := make(map[int]int)
	return helper(n, result)
}

func helper(n int, result map[int]int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	if v, ok := result[n]; ok {
		return v
	}

	v := helper(n-1, result) + helper(n-2, result)
	result[n] = v
	return v
}

// 方法四
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2

	}

	var v, v1, v2 int
	v1 = 2
	v2 = 1
	for i := 3; i <= n; i++ {
		v = v1 + v2
		v1, v2 = v, v1
	}
	return v
}
