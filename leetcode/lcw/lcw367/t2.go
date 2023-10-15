package lcw367

import "strings"

func shortestBeautifulSubstring(s string, k int) string {
	var check = func(t string) bool {
		var cnt = 0
		for _, v := range t {
			if v == '1' {
				cnt++
			}
		}
		return cnt == k
	}
	var n = len(s)
	var INF = strings.Repeat("1", n+1)
	var res = INF
	var valid = false
	for l := 1; l <= n; l++ {
		if valid {
			break
		}
		for i := 0; i <= n-l; i++ {
			var j = i + l - 1
			var t = s[i : j+1]
			if check(t) && t < res {
				res = t
				valid = true
			}
		}
	}
	if res == INF {
		return ""
	}
	return res
}
