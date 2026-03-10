func largestAltitude(gain []int) int {
    n := len(gain)
    prefGain := make([]int, n+1)
    for i := 0; i < len(gain); i++ {
        prefGain[i+1] = prefGain[i] + gain[i]
    }

    maxGain := prefGain[0]
    for _, v := range prefGain {
        if v > maxGain {
            maxGain = v
        }
    }

    return maxGain
}