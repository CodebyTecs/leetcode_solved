func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]byte][]string, len(strs))

	for _, s := range strs {
		var key [26]byte
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		groups[key] = append(groups[key], s)
	}

	res := make([][]string, 0, len(groups))
	for _, g := range groups {
		res = append(res, g)
	}
	return res
}