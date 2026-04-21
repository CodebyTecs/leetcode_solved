func transpose(matrix [][]int) [][]int {
    m := len(matrix)
    n := len(matrix[0])

    res := make([][]int, n)
    for i := range res {
        res[i] = make([]int, m)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            res[j][i] = matrix[i][j]
        }
    }

    return res
}