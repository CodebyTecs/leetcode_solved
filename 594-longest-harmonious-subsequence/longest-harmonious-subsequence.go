func findLHS(nums []int) int {
    freq := make(map[int]int)
    for _, x := range nums {
        freq[x]++
    }

    ans := 0
    for x, c := range freq {
        if freq[x+1] > 0 {
            ans = max(ans, c + freq[x+1])
        }
    }
    
    return ans
}