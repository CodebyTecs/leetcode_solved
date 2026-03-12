func lengthOfLongestSubstring(s string) int {
    maxLength := 0
    charCount := make(map[byte]int)
    l := 0

    for r := 0; r < len(s); r++ {
       charCount[s[r]]++
       for charCount[s[r]] > 1 {
            charCount[s[l]]--
            l++
        }
        maxLength = max(maxLength, r-l+1)
    } 

    return maxLength
}