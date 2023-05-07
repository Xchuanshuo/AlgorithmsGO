package test1625

import "strings"

func findLexSmallestString(s string, a int, b int) string {
	var add = func(s string) string {
		var bs = []byte(s)
		for i := 0; i < len(bs); i++ {
			if i%2 == 1 {
				bs[i] = byte('0' + (int(bs[i]-'0')+a)%10)
			}
		}
		return string(bs)
	}
	var flip = func(s string) string {
		var n = len(s)
		return s[n-b:] + s[0:n-b]
	}
	var q = []string{s}
	var visited = make(map[string]bool)
	var res = s
	for len(q) > 0 {
		var cur = q[0]
		if strings.Compare(cur, res) < 0 {
			res = cur
		}
		q = q[1:]
		var t1 = add(cur)
		if !visited[t1] && strings.Compare(t1, cur) < 0 {
			visited[t1] = true
			q = append(q, t1)
			res = cur
		}
		var t2 = flip(cur)
		if !visited[t2] && strings.Compare(t2, cur) < 0 {
			visited[t2] = true
			q = append(q, t2)
		}
	}
	return res
}
