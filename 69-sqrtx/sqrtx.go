func mySqrt(x int) int {
    left := 0
    right := x

    if x == 0 {
        return 0
    }

    for left <= right {
        n := (left + right) / 2
        if n*n <= x && (n+1)*(n+1) > x {
            return n
        } else if n*n > x {
            right = n - 1
        } else {
            left = n + 1
        }
    }

    return 0
}