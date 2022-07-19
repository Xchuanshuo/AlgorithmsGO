package test1021

import "strings"

func removeOuterParentheses(s string) string {
	var sb strings.Builder
	var lastCnt = 0
	cnt, idx := 0, 0
	for i, v := range s {
		if v == '(' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 1 && lastCnt == 0 {
			idx = i
		}
		if cnt > 0 && idx != i {
			sb.WriteRune(v)
		}
		lastCnt = cnt
	}
	return sb.String()
}
