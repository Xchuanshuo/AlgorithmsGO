package main

func fraction(cont []int) []int {
	var n = len(cont)
	var dfs func(idx int) []int
	dfs = func(idx int) []int {
		if idx == n-1 {
			return []int{cont[idx], 1}
		}
		var pre = dfs(idx + 1)
		pre[0], pre[1] = pre[1], pre[0]
		return []int{pre[1]*cont[idx] + pre[0], pre[1]}
	}
	return dfs(0)
}
