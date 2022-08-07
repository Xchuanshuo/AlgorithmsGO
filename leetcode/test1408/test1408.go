package test1408

import "strings"

func stringMatching(words []string) []string {
	var res = make([]string, 0)
	for i, w := range words {
		for j, other := range words {
			if i == j {
				continue
			}
			if strings.Index(other, w) != -1 {
				res = append(res, w)
				break
			}
		}
	}
	return res
}
