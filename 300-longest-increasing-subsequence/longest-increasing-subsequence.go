func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    dp := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        dp[i] = 1
    }

    res := 1
    for i := 0; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] && dp[i] < dp[j] + 1 {
                dp[i] = dp[j] + 1
            }
        }
        res = max(dp[i], res)
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}