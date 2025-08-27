func sortedSquares(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    l := 0
    r := n - 1
    i := n - 1

    for l <= r {
        left := nums[l]
        right := nums[r]
        if left*left > right*right {
            res[i] = left * left
            l++
        } else {
            res[i] = right * right
            r--
        }
        i--
    }
    return res
}