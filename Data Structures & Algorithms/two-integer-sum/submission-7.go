func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, val := range nums {
        diff := target - val
        if idx, ok := seen[diff]; ok {
            return []int{idx, i} 
        }
        seen[val] = i
    }
    return nil
}
