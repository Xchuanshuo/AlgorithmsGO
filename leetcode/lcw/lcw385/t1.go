package lcw385

import "strings"

func countPrefixSuffixPairs1(words []string) int {
	var isPrefixAndSuffix = func(s1, s2 string) bool {
		return strings.HasPrefix(s2, s1) && strings.HasSuffix(s2, s1)
	}
	var n = len(words)
	var res = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				res++
			}
		}
	}
	return res
}
