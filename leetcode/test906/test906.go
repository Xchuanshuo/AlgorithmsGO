package test906

import (
	"fmt"
	"strconv"
)

/**
 * @Description https://leetcode.cn/problems/super-palindromes/
 * idea: 数据范围1e18, 满足sqrt(num)后也为回文数, 枚举范围缩小为1e9, 还是会超时，
		 而求解对于不同数位长度的回文数，只需要枚举回文根，即数位长度/2的数即可，此时枚举范围缩小到1e5
		 使用循环或者递归构造回文数，最后对每个回文数的合法性进行判定即可
 **/

var src = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func superpalindromesInRange(left string, right string) int {
	var calc = func(l int) []string {
		var dfs func(n int) []string
		dfs = func(n int) []string {
			if n == 0 {
				return []string{""}
			}
			if n == 1 {
				return src
			}
			var res = make([]string, 0)
			for _, s := range dfs(n - 2) {
				for x := 1; x <= 9; x++ {
					res = append(res, fmt.Sprintf("%d%s%d", x, s, x))
				}
				if n != l {
					res = append(res, "0"+s+"0")
				}
			}
			return res
		}
		return dfs(l)
	}
	var isPal = func(v int64) bool {
		var s = fmt.Sprintf("%d", v)
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-i-1] {
				return false
			}
		}
		return true
	}
	lv, _ := strconv.ParseInt(left, 10, 64)
	rv, _ := strconv.ParseInt(right, 10, 64)
	var res = 0
	for l := max(1, len(left)/2); l <= (len(right)+1)/2; l++ {
		for _, s := range calc(l) {
			curV, _ := strconv.ParseInt(s, 10, 64)
			curV *= curV
			if curV != 0 && curV >= lv && curV <= rv && isPal(curV) {
				res++
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
