func getRow(rowIndex int) []int {
    triangle := make([][]int, rowIndex + 1)

    for i := 0; i <= rowIndex; i++ {
        triangle[i] = make([]int, i + 1)
        triangle[i][0] = 1
        triangle[i][i] = 1
    }

    for i := 2; i <= rowIndex; i++ {
        for j := 1; j < i; j ++ {
            triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
        }
    }

    return triangle[rowIndex]

}