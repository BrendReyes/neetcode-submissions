func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	max := 0

	for left < right {
		soFarMax := (right - left) * min(heights[left], heights[right])
		if soFarMax > max {
			max = soFarMax
		}

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return max
}
