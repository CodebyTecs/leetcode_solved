func pivotIndex(nums []int) int {
    n := len(nums)
    prefixRight := make([]int, n+1)
    prefixLeft := make([]int, n+1)

    for i := 0; i < n; i++ {
        prefixRight[i+1] = prefixRight[i] + nums[n-i-1]
    }

    for i := 0; i < n; i++ {
        prefixLeft[i+1] = prefixLeft[i] + nums[i]
    }

    for i := 0; i < n; i++ {
        if prefixLeft[i] == prefixRight[n-i-1] {
            return i
        }
    }

    return -1
}