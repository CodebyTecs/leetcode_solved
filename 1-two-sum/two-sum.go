func twoSum(nums []int, target int) []int {
    hash := make(map[int]int, len(nums))
    for i, v := range nums {
        hash[v] = i
    }
    
    for i, n := range nums {
        j := target - n
        if i != hash[j] {
            _, ok := hash[j]
            if ok {
                return []int{i, hash[j]}
            }
        }
    }
    
    return nil
}