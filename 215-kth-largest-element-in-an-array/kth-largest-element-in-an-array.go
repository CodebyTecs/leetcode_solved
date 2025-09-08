func findKthLargest(nums []int, k int) int {
    n := len(nums)
    target := n - k

    l, r := 0, n-1
    for l <= r {
        p := partition(nums, l, r)
        if p == target {
            return nums[p]
        } else if p < target {
            l = p + 1
        } else {
            r = p - 1
        }
    }
    return -1
}

func partition(a []int, l, r int) int {
    pivot := a[r]
    i := l
    for j := l; j < r; j++ {
        if a[j] < pivot {
            a[i], a[j] = a[j], a[i]
            i++
        }
    }
    a[i], a[r] = a[r], a[i]
    return i
}