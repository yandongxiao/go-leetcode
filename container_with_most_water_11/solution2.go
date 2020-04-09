package container_with_most_water_11

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		h := min(height[left], height[right])
		area := h * (right - left)
		maxArea = max(maxArea, area)

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
