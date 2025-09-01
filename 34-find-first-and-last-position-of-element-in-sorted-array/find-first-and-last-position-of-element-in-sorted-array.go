func searchRange(nums []int, target int) []int {
    left := findLeft(nums, target)
    right := findRight(nums, target)
    return []int{left, right}
}

func findLeft(nums []int, target int) int {
    l, r := 0, len(nums)-1
    res := -1
    for l <= r {
        mid := (l + r) / 2
        if nums[mid] >= target {
            if nums[mid] == target {
                res = mid
            }
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return res
}

func findRight(nums []int, target int) int {
    l, r := 0, len(nums)-1
    res := -1
    for l <= r {
        mid := (l + r) / 2
        if nums[mid] <= target {
            if nums[mid] == target {
                res = mid
            }
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return res
}