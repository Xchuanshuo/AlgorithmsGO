package test953

func isAlienSorted(words []string, order string) bool {
	var mp = make(map[byte]int)
	for i := 0; i < len(order); i++ {
		mp[order[i]] = i
	}
	for i := 1; i < len(words); i++ {
		n1, n2 := len(words[i-1]), len(words[i])
		var n = min(n1, n2)
		var isValid = false
		for j := 0; j < n; j++ {
			if mp[words[i-1][j]] > mp[words[i][j]] {
				return false
			}
			if mp[words[i-1][j]] < mp[words[i][j]] {
				isValid = true
				break
			}
		}
		if !isValid && n1 > n2 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
