package lcw385

func mostFrequentPrime(g [][]int) int {
	m, n := len(g), len(g[0])
	var cnt = make(map[int]int)
	var cal = func(x, y int, dx, dy int) {
		var dfs func(x, y int, pre int)
		dfs = func(x, y int, pre int) {
			if x < 0 || x >= m || y < 0 || y >= n {
				return
			}
			var cur = pre*10 + g[x][y]
			if cur > 10 && isPrim(cur) {
				cnt[cur]++
			}
			dfs(x+dx, y+dy, cur)
		}
		dfs(x, y, 0)
	}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for x := 0; x < m; x++ {
				for y := 0; y < n; y++ {
					cal(x, y, i, j)
				}
			}
		}
	}
	if len(cnt) == 0 {
		return -1
	}
	var res = -1
	for k, v := range cnt {
		if res == -1 || v > cnt[res] || (v == cnt[res] && k > res) {
			res = k
		}
	}
	return res
}

func isPrim(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
