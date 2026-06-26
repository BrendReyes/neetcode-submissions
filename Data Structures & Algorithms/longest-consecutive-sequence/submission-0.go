func longestConsecutive(nums []int) int {
	seen := make(map[int]bool)
	max := 0

	for _, num := range nums {
		if ok := seen[num]; !ok {
			seen[num] = true
		}
	}

	for _, num := range nums {
		if !seen[num - 1] {
			currentNum := num
            currentLen := 1

			for seen[currentNum+1] {
                currentNum++
                currentLen++
            }

            if currentLen > max {
                max = currentLen
            }
		}
	}

	return max
}
