func singleNumber(nums []int) int {
    count := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        count[nums[i]] += 1
    }

    for num, k := range count {
        if k == 1 {
            return num
        }
    }

    return 0
}