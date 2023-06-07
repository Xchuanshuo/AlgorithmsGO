package test902

import "strconv"

// 数位dp dfs

func atMostNGivenDigitSet(digits []string, n int) int {
	var s = strconv.Itoa(n)
	var l = len(s)
	var mem = make([]int, l)
	for i := 0; i < l; i++ {
		mem[i] = -1
	}
	var dfs func(i int, limit, isNum bool) int
	dfs = func(i int, limit, isNum bool) int {
		if i == l {
			if isNum {
				return 1
			}
			return 0
		}
		if !limit && isNum && mem[i] >= 0 {
			return mem[i]
		}
		var res = 0
		var up = 9
		if limit {
			up = int(s[i] - '0')
		}
		if !isNum {
			res += dfs(i+1, false, false)
		}
		for _, v := range digits {
			x := int(v[0] - '0')
			if x > up {
				break
			}
			res += dfs(i+1, limit && x == up, isNum || x > 0)
		}
		if !limit && isNum {
			mem[i] = res
		}
		return res
	}
	return dfs(0, true, false)
}
