package test890

func findAndReplacePattern(words []string, pattern string) []string {
	var res = make([]string, 0)
	for _, w := range words {
		var mapTable = make(map[byte]byte)
		var used = make(map[byte]bool)
		var isValid = true
		for i := 0; i < len(w); i++ {
			if v, ok := mapTable[w[i]]; ok {
				if v != pattern[i] {
					isValid = false
					break
				}
			} else {
				if used[pattern[i]] {
					isValid = false
					break
				}
				mapTable[w[i]] = pattern[i]
				used[pattern[i]] = true
			}
		}
		if isValid {
			res = append(res, w)
		}
	}
	return res
}
