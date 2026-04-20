func maxArea(height []int) int {
    maxWater := 0

    l := 0
    r := len(height) - 1

    for l <= r {
        countWater := min(height[l], height[r]) * (r-l)
        if countWater > maxWater {
            maxWater = countWater
        }
        
        if height[l] <= height[r] {
            l++
        } else if height[l] > height[r] {
            r--
        }
    }

    return maxWater
}

func min(a,b int) int {
    if a < b {
        return a
    }

    return b
}