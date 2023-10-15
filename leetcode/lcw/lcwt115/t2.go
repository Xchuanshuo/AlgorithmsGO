package lcwt115

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	var cal = func(f int) []string {
		var res = make([]string, 0)
		var pre = f
		for i, v := range groups {
			if v != pre {
				res = append(res, words[i])
				pre = v
			}
		}
		return res
	}
	var res1 = cal(0)
	var res2 = cal(1)
	if len(res1) > len(res2) {
		return res1
	}
	return res2
}
