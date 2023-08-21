package lcwt111

// 数位dp mem[i][mod][sum]表示前i位数，%k值为mod, 且奇数和偶数差为sum的条件下 总的方案数

import "strconv"

func numberOfBeautifulIntegers(low int, high int, k int) int {
	var calc = func(num string) int {
		// 表示从i位数开始，初始数位和为sum的情况下的整数个数
		var dfs func(i int, mod, sum int, limit, isNum bool) int
		var mem = make([][]map[int]int, len(num))
		for i := 0; i < len(mem); i++ {
			mem[i] = make([]map[int]int, k)
			for j := 0; j < len(mem[i]); j++ {
				mem[i][j] = make(map[int]int)
			}
		}
		dfs = func(i int, mod, sum int, limit, isNum bool) int {
			if i >= len(num) {
				if sum == 0 && mod == 0 && isNum {
					return 1
				}
				return 0
			}
			if v, exist := mem[i][mod][sum]; exist && !limit && isNum {
				return v
			}
			var up = int(num[i] - '0')
			if !limit {
				up = 9
			}
			var res = 0
			for x := 0; x <= up; x++ {
				var cur = sum
				if isNum || x > 0 { // 边界判定，防止前导0导致计入不合法的答案
					if x%2 == 0 {
						cur += 1
					} else {
						cur -= 1
					}
				}
				res += dfs(i+1, (mod*10+x)%k, cur, limit && x == up, isNum || x > 0)
			}
			if !limit && isNum {
				mem[i][mod][sum] = res
			}
			return res
		}
		return dfs(0, 0, 0, true, false)
	}
	return calc(strconv.Itoa(high)) - calc(strconv.Itoa(low-1))
}
