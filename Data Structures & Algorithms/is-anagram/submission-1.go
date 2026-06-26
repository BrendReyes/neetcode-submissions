import "maps"

func isAnagram(s string, t string) bool {
    exists1:= make(map[rune]int)
    exists2 := make(map[rune]int)

    for _, char := range s {
        exists1[char]++
    }

    for _, char := range t {
        exists2[char]++
    }

    return maps.Equal(exists1, exists2)
}
