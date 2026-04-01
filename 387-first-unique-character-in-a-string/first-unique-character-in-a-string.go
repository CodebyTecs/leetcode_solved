func firstUniqChar(s string) int {
    if len(s) == 1 {
        return 0
    }

    freqCh := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        freqCh[s[i]]++
    }

    for i := 0; i < len(s); i++ {
        if freqCh[s[i]] == 1 {
            return i
        }
    }

    return -1
}