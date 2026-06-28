func productExceptSelf(nums []int) []int {
	res := []int{}

	for i := 0; i < len(nums); i++ {
		product := 1
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			} else {
				product *= nums[j]
			}
			
		}
		res = append(res, product)
	}

	return res
}
