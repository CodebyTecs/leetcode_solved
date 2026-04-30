func combine(n int, k int) [][]int {
    res := make([][]int, 0)
    cur := make([]int, 0)

    backtrack(&res, &cur, n, k, 1)

    return res
}

func backtrack(res *[][]int, cur *[]int, n int, k int, start int) {
    if len(*cur) == k {
        comb := make([]int, len(*cur))
        copy(comb, *cur)
        *res = append(*res, comb)
        return
    }

    for i := start; i <= n; i++ {
        *cur = append(*cur, i)
        backtrack(res, cur, n, k, i+1)
        *cur = (*cur)[:len(*cur)-1]
    }
}