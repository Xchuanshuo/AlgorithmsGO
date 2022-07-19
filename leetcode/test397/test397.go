package test397

var mem = make(map[int]int)
func integerReplacement(n int) int {
	return dfs(n)
}

func dfs(n int) int {
	if n == 1 {
		return 0
	}
	if _, ok := mem[n]; ok {
		return mem[n]
	}
	cur := 0
	if (n&1) == 0 {
		cur = 1 + dfs(n / 2)
	} else {
		cur = 1 + min(dfs(n+1), dfs(n-1))
	}
	mem[n] = cur
	return cur
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}