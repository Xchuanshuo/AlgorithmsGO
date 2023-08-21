package test2193

import "strings"

func minMovesToMakePalindrome(s string) int {
	if len(s) <= 1 {
		return 0
	}
	var n = len(s)
	l, r := s[0], s[n-1]
	if l == r {
		return minMovesToMakePalindrome(s[1 : len(s)-1])
	}
	var ri = strings.LastIndex(s, string(l))
	var li = strings.Index(s, string(r))
	l1, l2 := li, n-ri-1
	if ri == -1 || l1 <= l2 {
		return l1 + minMovesToMakePalindrome(s[0:li]+s[li+1:n-1])
	}
	if li == -1 || l2 < l1 {
		return l2 + minMovesToMakePalindrome(s[1:ri]+s[ri+1:])
	}
	return 0
}
