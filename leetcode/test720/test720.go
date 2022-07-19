package test720

import "sort"

func longestWord(words []string) string {
	sort.Strings(words)
	var set = make(map[string]bool)
	var res = ""
	for _, w := range words {
		if len(w) == 1 || set[w[0:len(w)-1]] {
			if len(w) > len(res) {
				res = w
			}
			set[w] = true
		}
	}
	return res
}
