package test761

import (
	"sort"
	"strings"
)

func makeLargestSpecial(s string) string {
	i, cnt := 0, 0
	var list = make([]string, 0)
	for j, v := range s {
		if v == '1' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			list = append(list, "1"+makeLargestSpecial(s[i+1:j])+"0")
			i = j + 1
		}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})
	var sb strings.Builder
	for _, l := range list {
		sb.WriteString(l)
	}
	return sb.String()
}
