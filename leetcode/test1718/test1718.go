package test1718

// 回溯 + 贪心填数，依次填最大的，找到第一个合法的情况下返回答案。
//	数据范围20，可以使用位运算记录状态

func constructDistancedSequence(n int) []int {
	var l = 2*n - 1
	var res = make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = -1
	}
	var dfs func(idx int, visited int) bool
	dfs = func(idx int, visited int) bool {
		if idx == l {
			return true
		}
		if res[idx] != -1 {
			return dfs(idx+1, visited)
		}
		for i := n; i >= 1; i-- {
			if visited&(1<<i) != 0 || (i != 1 && (idx+i >= l || res[idx+i] != -1)) {
				continue
			}
			res[idx] = i
			if i != 1 {
				res[idx+i] = i
			}
			var valid = dfs(idx+1, visited|1<<i)
			if valid {
				return true
			}
			res[idx] = -1
			if i != 1 {
				res[idx+i] = -1
			}
		}
		return false
	}
	dfs(0, 0)
	return res
}
