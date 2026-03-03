func totalFruit(fruits []int) int {
    l := 0
    maxCount := 0
    fruitCount := make(map[int]int)

    for r := 0; r < len(fruits); r++ {
        fruitCount[fruits[r]]++

        for len(fruitCount) > 2 {
            fruitCount[fruits[l]]--

            if fruitCount[fruits[l]] == 0 {
                delete(fruitCount, fruits[l])
            }
            l++
        }

        maxCount = max(maxCount, r-l+1)
    }

    return maxCount
}