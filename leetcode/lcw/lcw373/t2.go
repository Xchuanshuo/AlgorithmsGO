package lcw373

import "strings"

func beautifulSubstrings1(s string, k int) int {
	var n = len(s)
	var res = 0
	for i := 0; i < n; i++ {
		v, c := 0, 0
		for j := i; j < n; j++ {
			if strings.ContainsAny("aeiou", string(s[j])) {
				v++
			} else {
				c++
			}
			if v == c && v*c%k == 0 {
				res++
			}
		}
	}
	return res
}
