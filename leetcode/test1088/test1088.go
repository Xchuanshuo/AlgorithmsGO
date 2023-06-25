package test1088

import "strconv"

var mp = map[int]int{1: 1, 6: 9, 8: 8, 9: 6, 0: 0}

func confusingNumberII(n int) int {
	var calc = func(num string) int {
		// 表示从i位数开始，初始数位和为sum的情况下的整数个数
		var dfs func(i int, val int, limit bool) int
		dfs = func(i int, val int, limit bool) int {
			if i >= len(num) {
				if isValid(val) {
					return 1
				}
				return 0
			}
			var up = int(num[i] - '0')
			if !limit {
				up = 9
			}
			var res = 0
			for j := 0; j <= up; j++ {
				if _, exist := mp[j]; !exist {
					continue
				}
				res += dfs(i+1, val*10+j, limit && j == up)
			}
			return res
		}
		return dfs(0, 0, true)
	}
	return calc(strconv.Itoa(n))
}

func isValid(n int) bool {
	var t = n
	var v = 0
	for t != 0 {
		if _, ok := mp[t%10]; !ok {
			return false
		}
		v = v*10 + mp[t%10]
		t /= 10
	}
	return v != n
}
