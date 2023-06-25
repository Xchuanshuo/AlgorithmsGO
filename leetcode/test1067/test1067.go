package test1067

import "strconv"

/**
 * @Description https://leetcode.cn/problems/digit-count-in-range/
 * idea: 数位dp模板 求[low,high]范围内所有的整数中 数位为d的总出现次数
 **/

func digitsCount(d int, low int, high int) int {
	var cal = func(d, n int) int {
		var s = strconv.Itoa(n)
		var l = len(s)
		var mem = make([][]int, l)
		for i := 0; i < l; i++ {
			mem[i] = make([]int, l)
			for j := 0; j < l; j++ {
				mem[i][j] = -1
			}
		}
		var dfs func(i, cntD int, limit, isNum bool) int
		dfs = func(i, cntD int, limit, isNum bool) int {
			if i == l {
				return cntD
			}
			if !limit && isNum && mem[i][cntD] >= 0 {
				return mem[i][cntD]
			}
			var up = 9
			if limit {
				up = int(s[i] - '0')
			}
			var res = 0
			for x := 0; x <= up; x++ {
				var add = 0
				if x == d && (d > 0 || (d == 0 && isNum)) {
					add = 1
				}
				res += dfs(i+1, cntD+add, limit && x == up, isNum || x > 0)
			}
			if !limit && isNum {
				mem[i][cntD] = res
			}
			return res
		}
		return dfs(0, 0, true, false)
	}
	return cal(d, high) - cal(d, low-1)
}
