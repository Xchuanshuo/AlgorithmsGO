package test796

import "strings"

func rotateString(s string, goal string) bool {
	var t = s + s
	return len(s) == len(goal) && strings.Contains(t, goal)
}
