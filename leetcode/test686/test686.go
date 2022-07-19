package test686

import "strings"

func repeatedStringMatch(a string, b string) int {
	la, lb := len(a), len(b)
	cnt := lb / la + 2
	var sb strings.Builder
	for i := 1;i <= cnt;i++ {
		sb.WriteString(a)
		if strings.Contains(sb.String(), b) {
			return i
		}
	}
	return -1
}