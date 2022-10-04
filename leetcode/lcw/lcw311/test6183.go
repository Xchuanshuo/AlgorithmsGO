package lcw311

func sumPrefixScores(words []string) []int {
	var scoreMap = make(map[string]int)
	for _, w := range words {
		for i := 1; i <= len(w); i++ {
			scoreMap[w[0:i]]++
		}
	}
	var res = make([]int, len(words))
	for idx, w := range words {
		var cur = 0
		for i := 1; i <= len(w); i++ {
			cur += scoreMap[w[0:i]]
		}
		res[idx] = cur
	}
	return res
}
