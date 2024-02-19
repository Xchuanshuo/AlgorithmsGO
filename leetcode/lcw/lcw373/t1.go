package lcw373

func areSimilar(g [][]int, k int) bool {
	m, n := len(g), len(g[0])
	for i := 0; i < m; i++ {
		var s = 0
		k %= n
		if i%2 == 0 { // 奇数行
			s = (n - k + n) % n
		} else {
			s = (n + k) % n
		}
		for j := 0; j < n; j++ {
			var move = (s + j) % n
			if g[i][j] != g[i][move] {
				return false
			}
		}
	}
	return true
}
