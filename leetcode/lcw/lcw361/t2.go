package lcw361

var INF = int(1e9)

func minimumOperations(num string) int {
	var n = len(num)
	var mem = make([][26]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < len(mem[i]); j++ {
			mem[i][j] = -1
		}
	}
	var dfs func(i, v int) int
	dfs = func(i, v int) int {
		if i == n {
			if v == 0 {
				return 0
			}
			return INF
		}
		if mem[i][v] != -1 {
			return mem[i][v]
		}
		var c = int(num[i] - '0')
		var cur = min(dfs(i+1, (v*10+c)%25), dfs(i+1, v)+1)
		mem[i][v] = cur
		return cur
	}
	return dfs(0, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
