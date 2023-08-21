package test1444

var M = int(1e9) + 7

func ways(grid []string, k int) int {
	var g = make([][]byte, 0)
	for _, v := range grid {
		g = append(g, []byte(v))
	}
	m, n := len(g), len(g[0])
	var sum = make([][]int, m+1)
	for i := 0; i < m/2; i++ {
		g[i], g[m-i-1] = g[m-i-1], g[i]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			g[i][j], g[i][n-j-1] = g[i][n-j-1], g[i][j]
		}
	}
	for i := 0; i < len(sum); i++ {
		sum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			var cur = 0
			if g[i-1][j-1] == 'A' {
				cur = 1
			}
			sum[i][j] = cur + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
		}
	}
	var mem = initMem(m, n, k)
	var dfs func(i, j, k int) int
	dfs = func(i, j, k int) int {
		if i == 0 || j == 0 {
			return 0
		}
		if k == 1 {
			if sum[i][j] == 0 {
				return 0
			}
			return 1
		}
		if mem[i][j][k] != -1 {
			return mem[i][j][k]
		}
		var res = 0
		for r := i - 1; r >= 1; r-- {
			if r+j < k {
				break
			}
			if sum[i][j]-sum[r][j] == 0 {
				continue
			}
			res = (res + dfs(r, j, k-1)) % M
		}
		for c := j - 1; c >= 1; c-- {
			if i+c < k {
				break
			}
			if sum[i][j]-sum[i][c] == 0 {
				continue
			}
			res = (res + dfs(i, c, k-1)) % M
		}
		mem[i][j][k] = res
		return res
	}
	return dfs(m, n, k)
}

var initMem = func(m, n, k int) [][][]int {
	var mem = make([][][]int, m+1)
	for i := 0; i < len(mem); i++ {
		mem[i] = make([][]int, n+1)
		for j := 0; j < len(mem[i]); j++ {
			mem[i][j] = make([]int, k+1)
			for k := 0; k < len(mem[i][j]); k++ {
				mem[i][j][k] = -1
			}
		}
	}
	return mem
}
