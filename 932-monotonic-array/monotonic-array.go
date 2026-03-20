func isMonotonic(nums []int) bool {
    if len(nums) == 1 {
        return true
    }

    isDec := false
    isInc := false
    for i := 1; i < len(nums); i++ {
        if nums[i-1] < nums[i] {
            isInc = true
        } else if nums[i-1] > nums[i] {
            isDec = true
        }
    }

    if isDec && isInc {
        return false
    }

    return true
}