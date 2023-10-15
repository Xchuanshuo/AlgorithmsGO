package lcw367

// 费马小定理计算最终结果 (a/b)%p = a*b^(p-2)%p 使用快速幂进行计算

var M = int64(12345)

func constructProductMatrix(g [][]int) [][]int {
	m, n := len(g), len(g[0])
	var total = int64(1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			total = (total*int64(g[i][j]) + M) % M
		}
	}
	var res = make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = int(total * fastPow(int64(g[i][j]), M-2) % M)
		}
	}
	return res
}

func fastPow(a, k int64) int64 {
	if k == 0 {
		return 1
	}
	a %= M
	if k%2 == 0 {
		val := fastPow(a, k/2) % M
		return val * val % M
	}
	return a * fastPow(a, k-1) % M
}
