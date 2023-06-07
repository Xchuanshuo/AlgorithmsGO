package lcw348

/**
 * @Description https://leetcode.cn/problems/count-of-integers/
 * idea: 数位dp模板
 **/

var M = int(1e9) + 7

func count(num1 string, num2 string, min_sum int, max_sum int) int {
	var calc = func(num string) int {
		// 表示从i位数开始，初始数位和为sum的情况下的整数个数
		var dfs func(i int, sum int, limit bool) int
		var mem = make([][]int, len(num))
		for i := 0; i < len(mem); i++ {
			mem[i] = make([]int, 401)
			for j := 0; j < len(mem[i]); j++ {
				mem[i][j] = -1
			}
		}
		dfs = func(i int, sum int, limit bool) int {
			if sum > max_sum {
				return 0
			}
			if i >= len(num) {
				if sum >= min_sum && sum <= max_sum {
					return 1
				}
				return 0
			}
			if !limit && mem[i][sum] != -1 {
				return mem[i][sum]
			}
			var up = int(num[i] - '0')
			if !limit {
				up = 9
			}
			var res = 0
			for j := 0; j <= up; j++ {
				res = (res + dfs(i+1, sum+j, limit && j == up)) % M
			}
			if !limit {
				mem[i][sum] = res
			}
			return res
		}
		return dfs(0, 0, true)
	}

	var v1 = calc(num1)
	var s = 0
	for _, v := range num1 {
		s += int(v - '0')
	}
	var add = 0
	if s >= min_sum && s <= max_sum {
		add = 1
	}
	var v2 = calc(num2)
	return (v2 - v1 + add + M) % M
}
