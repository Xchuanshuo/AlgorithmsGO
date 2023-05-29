package lcw346

import (
	"strconv"
)

var a []int

func init() {
	var dfs func(s string, t int) bool
	dfs = func(s string, t int) bool {
		if t == 0 && len(s) == 0 {
			return true
		}
		for i := 1; i <= len(s); i++ {
			v, _ := strconv.Atoi(s[0:i])
			if v <= t && dfs(s[i:], t-v) {
				return true
			}
		}
		return false
	}
	for i := 1; i <= 1000; i++ {
		if dfs(strconv.Itoa(i*i), i) {
			a = append(a, i)
		}
	}
}

func punishmentNumber(n int) int {
	var res = 0
	for _, v := range a {
		if v > n {
			break
		}
		res += v * v
	}
	return res
}
