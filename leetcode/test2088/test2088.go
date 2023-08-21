package test2088

func countPyramids(g [][]int) int {
	m, n := len(g), len(g[0])
	var res = 0
	var h = make([]int, n+2)
	var cal = func(i int) {
		var w = make([]int, n+2)
		for j := n; j >= 1; j-- {
			if g[i-1][j-1] == 0 {
				continue
			}
			w[j] = w[j+1] + 1
		}
		for j := 1; j <= n; j++ {
			if g[i-1][j-1] == 0 {
				h[j] = 0
				continue
			}
			var curH = (min(w[j-1], w[j+1])*2 + 1) / 2
			h[j] = min(h[j], curH) + 1
			if h[j] > 1 {
				res += h[j] - 1
			}
			w[j] = w[j-1] + 1
		}
	}
	for i := 1; i <= m; i++ {
		cal(i)
	}
	h = make([]int, n+2)
	for i := m; i >= 1; i-- {
		cal(i)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
