package lcw356

var M = int(1e9) + 7

func countSteppingNumbers(low string, high string) int {
	var calc = func(s string) int {
		var l = len(s)
		var mem = make([][]int, l)
		for i := 0; i < l; i++ {
			mem[i] = make([]int, 10)
			for j := 0; j < len(mem[i]); j++ {
				mem[i][j] = -1
			}
		}
		var dfs func(i, c int, limit, isNum bool) int
		dfs = func(i, c int, limit, isNum bool) int {
			if i == l {
				if !isNum {
					return 0
				}
				return 1
			}
			if !limit && isNum && mem[i][c] > 0 {
				return mem[i][c]
			}
			var up = 9
			if limit {
				up = int(s[i] - '0')
			}
			var res = 0
			for x := 0; x <= up; x++ {
				if !isNum {
					res = (res + dfs(i+1, x, limit && x == up, isNum || x > 0)) % M
				} else if abs(x-c) == 1 {
					res = (res + dfs(i+1, x, limit && x == up, isNum || x > 0)) % M
				}
			}
			if !limit && isNum {
				mem[i][c] = res
			}
			return res
		}
		return dfs(0, 0, true, false)
	}
	var add = 1
	for i := 1; i < len(low); i++ {
		if abs(int(low[i]-'0')-int(low[i-1]-'0')) != 1 {
			add = 0
			break
		}
	}
	return (calc(high) - calc(low) + add + M) % M
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
