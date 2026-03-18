func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	mp := map[int]int{0: 1}

	for _, v := range nums {
		sum += v
		count += mp[sum-k]
		mp[sum]++
	}

	return count
}