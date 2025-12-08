func isSubsequence(s string, t string) bool {
    left := len(s)
    right := len(t)

    pleft, pright := 0, 0

    for pleft < left && pright < right {
        if s[pleft] == t[pright] {
            pleft++
        }
        pright++
    }

    return pleft == left
}