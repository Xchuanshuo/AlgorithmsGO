package test1861

// 贪心 对于同一批翻转操作，顺序是【无关】的
// 	  1.翻转较高位的位能使答案更优, 所以考虑先进行行翻转将第一列全变为1
// 	  2.然后考虑列翻转, 将每列翻转更多1的个数
//   为什么先进行行翻转: 对于高位，若高位全为0，列翻转可以全变为1；若不全为0，列翻转只能将当前列0-1的数量互换，而行翻转可以将所有0变为1；
//					所以优先选择翻转行，可以包含两种情况。

func matrixScore(g [][]int) int {
	m, n := len(g), len(g[0])
	var res = m * (1 << (n - 1))
	for j := 1; j < n; j++ {
		var cnt0 = 0
		for i := 0; i < m; i++ {
			cnt0 += g[i][0] ^ g[i][j]
		}
		res += max(cnt0, m-cnt0) * (1 << (n - j - 1))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
