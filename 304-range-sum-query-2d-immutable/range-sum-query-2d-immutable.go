type NumMatrix struct {
    prefixMatrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return NumMatrix{}
    }

    rows, cols := len(matrix), len(matrix[0])
    prefixMatrix := make([][]int, rows)
    for i := range prefixMatrix {
        prefixMatrix[i] = make([]int, cols)
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            prefixMatrix[i][j] = matrix[i][j]
            if i > 0 {
                prefixMatrix[i][j] += prefixMatrix[i-1][j]
            }
            if j > 0 {
                prefixMatrix[i][j] += prefixMatrix[i][j-1]
            }
            if i > 0 && j > 0 {
                prefixMatrix[i][j] -= prefixMatrix[i-1][j-1]
            }
        }
    }

    return NumMatrix{prefixMatrix: prefixMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    res := this.prefixMatrix[row2][col2]
    if row1 > 0 {
        res -= this.prefixMatrix[row1-1][col2]
    }
    if col1 > 0 {
        res -= this.prefixMatrix[row2][col1-1]
    }
    if row1 > 0 && col1 > 0 {
        res += this.prefixMatrix[row1-1][col1-1]
    }
    return res
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */