package test709

import "strings"

func toLowerCase(s string) string {
	sb := strings.Builder{}
	for _, c := range s {
		t := c
		if t >= 'A' && t <= 'Z' {
			t = 'a' + t - 'A'
		}
		sb.WriteRune(t)
	}
	return sb.String()
}