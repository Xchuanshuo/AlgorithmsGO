package test2267

var dirs = [][]int{{1, 0}, {0, 1}}

func hasValidPath(g [][]byte) bool {
	m, n := len(g), len(g[0])
	if (m+n-1)%2 == 1 || g[0][0] == ')' || g[m-1][n-1] == '(' {
		return false
	}
	var mem = make([]map[int]bool, m+n)
	for i := 0; i < len(mem); i++ {
		mem[i] = make(map[int]bool)
	}
	var dfs func(x, y int, s []byte) bool
	dfs = func(x, y int, s []byte) bool {
		var p = x*n + y
		if v, exist := mem[len(s)][p]; exist {
			return v
		}
		if x == m-1 && y == n-1 {
			return len(s) == 0
		}
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= m || ny >= n {
				continue
			}
			if g[nx][ny] == ')' && len(s) == 0 {
				continue
			}
			if g[nx][ny] == ')' {
				s = s[0 : len(s)-1]
			} else {
				s = append(s, g[nx][ny])
			}
			if dfs(nx, ny, s) {
				return true
			}
			if g[nx][ny] == ')' {
				s = append(s, '(')
			} else {
				s = s[0 : len(s)-1]
			}
		}
		mem[len(s)][p] = false
		return false
	}
	return dfs(0, 0, []byte{'('})
}
