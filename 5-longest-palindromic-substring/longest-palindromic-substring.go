func longestPalindrome(s string) string {
    n := len(s)

    if n < 2 {
        return s
    }

    longestPal := make([]rune, 0, n)
    for i := 0; i < n; i++ {
        for j := i+1; j <= n; j++ {
            if isPalidrome(s[i:j]) {
                if len(longestPal) < len([]rune(s[i:j])) {
                    longestPal = []rune(s[i:j])
                }
            }
        }
    }

    return string(longestPal)
}

func isPalidrome(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-i-1] {
            return false
        }
    }

    return true
}