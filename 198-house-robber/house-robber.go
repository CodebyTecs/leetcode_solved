func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }

    dp1 := nums[0]
    dp2 := max(nums[0], nums[1])

    for i := 2; i < len(nums); i++ {
        dpNew := max(dp2, dp1 + nums[i])
        dp1 = dp2
        dp2 = dpNew
    }

    return dp2
}
func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}