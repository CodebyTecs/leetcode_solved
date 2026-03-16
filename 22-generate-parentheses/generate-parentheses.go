func generateParenthesis(n int) []string {
    res := make([]string, 0)
    backtrack(&res, n, 0, 0, "")
    return res
}

func backtrack(res *[]string, n, openN, closedN int, val string) {
    if openN == n && closedN == n {
        *res = append(*res, val)
        return
    }
    if openN < n {
        backtrack(res, n, openN+1, closedN, val+"(")
    }
    if closedN < openN {
        backtrack(res, n, openN, closedN+1, val+")")
    }
}