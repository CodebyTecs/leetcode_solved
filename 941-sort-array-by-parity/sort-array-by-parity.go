func sortArrayByParity(nums []int) []int {
    if len(nums) == 1 {
        return nums
    }

    l := 0
    r := len(nums)-1

    for l < r {
        if nums[l]%2 == 0 {
            l++
        } else if nums[r]%2 == 1 {
            r--
        } else {
            nums[l], nums[r] = nums[r], nums[l]
            l++
            r--
        }
    }

    return nums
}