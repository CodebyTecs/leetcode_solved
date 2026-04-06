func diagonalSum(mat [][]int) int {
    n := len(mat)
    res := 0

    for i := 0; i < n; i++ {
        if i != n-1-i {
            res += mat[i][i]
            res += mat[i][n-1-i]
        } else {
            res += mat[i][i]
        }
    }

    return res
}