func intersection(nums1 []int, nums2 []int) []int {
    elems := make(map[int]struct{})
    res := make([]int, 0)

    for _, v := range nums1 {
        elems[v] = struct{}{}
    }

    for _, v := range nums2 {
        _, ok := elems[v]
        if ok {
            res = append(res, v)
            delete(elems, v)
        }
    }

    return res
}