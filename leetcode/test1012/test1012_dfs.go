package test1012

import (
	"strconv"
)

// 数位dp
func numDupDigitsAtMostN(n int) int {
	var s = strconv.Itoa(n)
	var l = len(s)
	var mem = make([][]int, l)
	for i := 0; i < l; i++ {
		mem[i] = make([]int, 1024)
		for j := 0; j < l; j++ {
			mem[i][j] = -1
		}
	}
	var dfs func(i, mask int, limit, isNum bool) int
	dfs = func(i, mask int, limit, isNum bool) int {
		if i == l {
			if isNum {
				return 1
			}
			return 0
		}
		if !limit && isNum && mem[i][mask] > 0 {
			return mem[i][mask]
		}
		var up = 9
		if limit {
			up = int(s[i] - '0')
		}
		var res = 0
		for x := 0; x <= up; x++ {
			if (1<<x)&mask > 0 {
				continue
			}
			var nxt = mask | 1<<x
			if !isNum && x == 0 {
				nxt = mask
			}
			res += dfs(i+1, nxt, limit && x == up, isNum || x > 0)
		}
		if !limit && isNum {
			mem[i][mask] = res
		}
		return res
	}
	return n - dfs(0, 0, true, false)
}
