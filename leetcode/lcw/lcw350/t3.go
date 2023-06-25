package lcw350

var M = int(1e9) + 7

func specialPerm(nums []int) int {
	var n = len(nums)
	var l = (1 << n) - 1
	var mem = make([][]int, n)
	for i := 0; i < n; i++ {
		mem[i] = make([]int, l+1)
		for j := 0; j < len(mem[i]); j++ {
			mem[i][j] = -1
		}
	}
	var dfs func(pre, s int) int
	dfs = func(pre, s int) int {
		if s == l {
			return 1
		}
		if pre >= 0 && mem[pre][s] != -1 {
			return mem[pre][s]
		}
		var res = 0
		for i := 0; i < n; i++ {
			if s&(1<<i) > 0 || (pre != -1 && nums[i]%nums[pre] != 0 && nums[pre]%nums[i] != 0) {
				continue
			}
			res = (res + dfs(i, s|1<<i)) % M
		}
		if pre >= 0 {
			mem[pre][s] = res
		}
		return res
	}
	return dfs(-1, 0)
}
