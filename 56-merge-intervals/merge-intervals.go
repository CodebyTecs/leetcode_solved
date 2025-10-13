func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }

    sort.Slice(intervals, func(i int, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    res := make([][]int, 0, len(intervals))
    curStart, curEnd := intervals[0][0], intervals[0][1]

    for i := 0; i < len(intervals); i++ {
        s, e := intervals[i][0], intervals[i][1]
        if s <= curEnd {
            if e > curEnd {
                curEnd = e
            }
        } else {
            res = append(res, []int{curStart, curEnd})
            curStart, curEnd = s, e
        }
    }

    res = append(res, []int{curStart, curEnd})
    return res
}