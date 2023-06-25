package test1397

import "strings"

/**
 * @Description https://leetcode.cn/problems/find-all-good-strings/
 * idea: 数位dp + kmp
 **/

var Mod = int(1e9) + 7

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	var next = getNext(evil)
	var calc = func(s string) int {
		var l = len(s)
		var mem = make([][]int, l)
		for i := 0; i < l; i++ {
			mem[i] = make([]int, len(evil))
			for j := 0; j < len(mem[i]); j++ {
				mem[i][j] = -1
			}
		}
		var dfs func(i, j int, limit, isNum bool) int
		dfs = func(i, j int, limit, isNum bool) int {
			if j == len(evil) {
				return 0
			}
			if i == l {
				return 1
			}
			if !limit && isNum && mem[i][j] > 0 {
				return mem[i][j]
			}
			var up = 25
			if limit {
				up = int(s[i] - 'a')
			}
			var res = 0
			for k := 0; k <= up; k++ {
				x := byte('a' + k)
				nj := j
				for nj > 0 && x != evil[nj] {
					nj = next[nj-1] + 1
				}
				if x == evil[nj] {
					nj++
				}
				res = (res + dfs(i+1, nj, limit && k == up, isNum || x > 0)) % Mod
			}
			if !limit && isNum {
				mem[i][j] = res
			}
			return res
		}
		return dfs(0, 0, true, false)
	}
	var add = 0
	if !strings.Contains(s1, evil) {
		add = 1
	}
	return (calc(s2) - calc(s1) + add + Mod) % Mod
}

func getNext(pattern string) []int {
	next := make([]int, len(pattern))
	next[0] = -1
	k := -1
	for i := 1; i < len(pattern); i++ {
		for k != -1 && pattern[k+1] != pattern[i] {
			k = next[k]
		}
		if pattern[k+1] == pattern[i] {
			k++
		}
		next[i] = k
	}
	return next
}
