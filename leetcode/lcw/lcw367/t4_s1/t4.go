package t4

// 前缀和

var M = 12345

func constructProductMatrix(g [][]int) [][]int {
	m, n := len(g), len(g[0])
	var pm = make([][]int, m)
	var sm = make([][]int, m)
	var pmr = make([]int, m)
	var smr = make([]int, m)
	for i := 0; i < m; i++ {
		pm[i] = make([]int, n)
		sm[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		var cur = 1
		for j := 0; j < n; j++ {
			cur = (cur*g[i][j] + M) % M
			pm[i][j] = cur
		}
		cur = 1
		for j := n - 1; j >= 0; j-- {
			cur = (cur*g[i][j] + M) % M
			sm[i][j] = cur
		}
		var pre = 1
		if i >= 1 {
			pre = pmr[i-1]
		}
		pmr[i] = (pre * pm[i][n-1]) % M
	}
	for i := m - 1; i >= 0; i-- {
		var pre = 1
		if i+1 < m {
			pre = smr[i+1]
		}
		smr[i] = (pre * pm[i][n-1]) % M
	}
	var res = make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pre, suf := 1, 1
			cpre, csuf := 1, 1
			if i-1 >= 0 {
				pre = pmr[i-1]
			}
			if i+1 < m {
				suf = smr[i+1]
			}
			if j-1 >= 0 {
				cpre = pm[i][j-1]
			}
			if j+1 < n {
				csuf = sm[i][j+1]
			}
			res[i][j] = ((pre * suf) % M * (cpre * csuf) % M) % M
		}
	}
	return res
}
