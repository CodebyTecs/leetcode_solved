func convertToTitle(columnNumber int) string {
    res := []rune{}

    for columnNumber > 0 {
        columnNumber--
        char := rune('A' + (columnNumber % 26))
        res = append(res, char)
        columnNumber /= 26
    }

    for i := 0; i < len(res)/2; i++ {
        res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
    }

    return string(res)
}