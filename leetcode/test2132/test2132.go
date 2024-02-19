package test2132

/**
 * @Description https://leetcode.cn/problems/stamping-the-grid/
 * idea: 贪心 + 二维差分矩阵 + 矩阵前缀和
		1.格子覆盖多次无所谓，若某个>=指定w和h的矩阵【和为0】，则说明以该处为右下角可以贴一枚邮票
		 贴邮票使用差分矩阵，O(1)时间更新4个位置覆盖邮票数量即可 (左上角, 右下角)+1 (右上界，左下角)-1
	    2.最终统计所有位置，若贴邮票数量=0, 说明无法贴满，返回false即可，否则可以贴满
 **/

func possibleToStamp(g [][]int, h int, w int) bool {
	m, n := len(g), len(g[0])
	var s = make([][]int, m+1)
	var dif = make([][]int, m+2)
	for i := 0; i < len(s); i++ {
		s[i] = make([]int, n+1)
	}
	for i := 0; i < len(dif); i++ {
		dif[i] = make([]int, n+2)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + g[i-1][j-1]
		}
	}
	for i := h; i <= m; i++ {
		for j := w; j <= n; j++ {
			var val = s[i][j] - s[i-h][j] - s[i][j-w] + s[i-h][j-w]
			if val == 0 { // 差分矩阵计数
				dif[i-h+1][j-w+1]++
				dif[i+1][j+1]++
				dif[i-h+1][j+1]--
				dif[i+1][j-w+1]--
			}
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 差分矩阵数值还原
			dif[i][j] += dif[i-1][j] + dif[i][j-1] - dif[i-1][j-1]
			if dif[i][j] == 0 && g[i-1][j-1] == 0 {
				return false
			}
		}
	}
	return true
}
