package test884

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	var set = make(map[string]int)
	for _, s := range strings.Fields(s1) { set[s]++ }
	for _, s := range strings.Fields(s2) { set[s]++ }
	res := make([]string, 0)
	for k, v := range set {
		if v == 1 { res = append(res, k) }
	}
	return res
}