package test248

/**
 * @Description https://leetcode.cn/problems/strobogrammatic-number-iii/
 * idea: dfs, 依次构造数位为n的中心对称数，构造逻辑 对构造出来的子结果左右添加pairs数对
			  构造子结果数位+2的中心对称数
 **/

import (
	"fmt"
	"strconv"
)

var pairs = [][]string{{"1", "1"}, {"6", "9"}, {"9", "6"}, {"8", "8"}}

func strobogrammaticInRange(low string, high string) int {
	var cal = func(n int) []string {
		var dfs func(x int) []string
		dfs = func(x int) []string {
			if x == 0 {
				return []string{""}
			}
			if x == 1 {
				return []string{"0", "1", "8"}
			}
			var res = make([]string, 0)
			for _, s := range dfs(x - 2) {
				for _, p := range pairs {
					res = append(res, fmt.Sprintf("%s%s%s", p[0], s, p[1]))
				}
				if x != n {
					res = append(res, "0"+s+"0")
				}
			}
			return res
		}
		return dfs(n)
	}
	lowV, _ := strconv.ParseInt(low, 10, 64)
	highV, _ := strconv.ParseInt(high, 10, 64)
	var res = 0
	for i := len(low); i <= len(high); i++ {
		for _, s := range cal(i) {
			cur, _ := strconv.ParseInt(s, 10, 64)
			if cur >= lowV && cur <= highV {
				res++
			}
		}
	}
	return res
}
