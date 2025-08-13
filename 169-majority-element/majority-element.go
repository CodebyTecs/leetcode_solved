func majorityElement(nums []int) int {
    count, candidate := 0, 0

    for _, num := range nums {
        if count == 0 {
            candidate = num
            count = 1
        } else if num == candidate {
            count++
        } else {
            count--
        }
    }

    return candidate
}