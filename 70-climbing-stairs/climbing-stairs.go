func climbStairs(n int) int {
    a := 1
    b := 1

    for i := 2; i <= n; i++ {
        a, b = a + b, a
    }

    return a
}