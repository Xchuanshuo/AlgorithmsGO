package lcwt118

import "strings"

func findWordsContaining(words []string, x byte) []int {
	var res = make([]int, 0)
	var t = string(x)
	for i, v := range words {
		if strings.Contains(v, t) {
			res = append(res, i)
		}
	}
	return res
}
