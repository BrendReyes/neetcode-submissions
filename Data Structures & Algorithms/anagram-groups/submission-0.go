func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)

    for _, str := range strs {
        var count [26]int
        for i, _ := range str {
            count[str[i]-'a']++
        }
        groups[count] = append(groups[count], str)
    }

    result := [][]string{}
    for _, group := range groups {
        result = append(result, group)
    }

    return result
}
