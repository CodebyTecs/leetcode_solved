func isPerfectSquare(num int) bool {
    if num == 1 {
        return true
    }
    
    l := 0
    r := num-1
    for l <= r {
        mid := (l+r)/2
        tempNum := mid*mid
        if tempNum == num {
            return true
        } else if tempNum > num {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }

    return false
}