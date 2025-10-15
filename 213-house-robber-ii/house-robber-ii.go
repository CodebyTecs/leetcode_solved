func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    if len(nums) == 2 {
        if nums[0] > nums[1] {
            return nums[0]
        } else {
            return nums[1]
        }
    }

    dp1 := nums[0]
    dp2 := max(nums[0], nums[1])
    for i := 2; i < len(nums) - 1; i++ {
        dpNew := max(dp2, dp1 + nums[i])
        dp1 = dp2
        dp2 = dpNew
    }

    dp3 := nums[1]
    dp4 := max(nums[1], nums[2])
    for i := 3; i < len(nums); i++ {
        dpNew := max(dp4, dp3 + nums[i])
        dp3 = dp4
        dp4 = dpNew
    }

    return max(dp2, dp4)
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}