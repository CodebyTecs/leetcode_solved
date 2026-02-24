func twoSum(numbers []int, target int) []int {
    n := len(numbers)
    res := make([]int, 2)
    l := 0
    r := n-1

    for l < r {
        if numbers[l] + numbers[r] == target {
            res[0] = l+1
            res[1] = r+1
            break
        } else if numbers[l] + numbers[r] > target {
            r--
        } else {
            l++
        }
    }

    return res
}