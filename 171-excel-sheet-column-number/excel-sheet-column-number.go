func titleToNumber(columnTitle string) int {
    res := 0

    for _, char := range columnTitle {
        val := int(char - 'A') + 1
        res = res*26 + val
    }

    return res
}