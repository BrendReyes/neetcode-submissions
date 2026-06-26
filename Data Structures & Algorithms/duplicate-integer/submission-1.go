func hasDuplicate(nums []int) bool {
    dups := make(map[int]bool)

    for _, val := range nums {
        _, ok := dups[val] 
        if ok {
            return true
        } 
        dups[val] = true
    }
    return false
}
