func missingNumber(nums []int) int {
    count := make(map[int]int)

    for i := 0; i <= len(nums); i++ {
        count[i] = 0
    }

    for _, key := range nums {
        count[key]++
    }

    for key, value := range count {
        if value == 0 {
            return key
        }
    }

    return -1
}