func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	buckets := make([][]int, len(nums) + 1)
	res := []int{}
	
	for _, num := range nums {
		freq[num]++
	}

	for key, val := range freq {
		buckets[val] = append(buckets[val], key)
	}

	for i := len(buckets) - 1; i >= 0; i-- {
        for _, num := range buckets[i] {
            res = append(res, num)
            
            if len(res) == k {
                return res
            }
        }
    }

	return res
}
