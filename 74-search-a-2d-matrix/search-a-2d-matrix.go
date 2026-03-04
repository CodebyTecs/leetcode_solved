func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }

    rows := len(matrix)
    cols := len(matrix[0])

    l := 0
    r := rows*cols - 1

    for l <= r {
        mid := l + (r-l)/2
        val := matrix[mid/cols][mid%cols]

        if val == target {
            return true
        } else if val > target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }

    return false
}