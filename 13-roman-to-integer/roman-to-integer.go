func romanToInt(s string) int {
    nums := map[byte]int {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    total := 0
    for i := 0; i < len(s); i++ {
        if i + 1 < len(s) && nums[s[i]] < nums[s[i + 1]] {
            total -= nums[s[i]]
        } else {
            total += nums[s[i]]
        }
    }

    return total
}