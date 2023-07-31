package test2245

/**
 * @Description https://leetcode.cn/problems/maximum-trailing-zeros-in-a-cornered-path/
 * idea:  1.路径乘积尾随0 => 路径上的数因子能分解为多少个2和5, 记录数量
		  2.至多可以拐一次弯 => 拐弯后的数即使不存在因子2和5，对结果也无影响
		  3.前缀和记录从上到下，从左至右的因子2、5的个数，对于每个位置考虑4种到达边界的路径
		    左上、右上、左下，右下，依次根据2和5的因子的最小数即为乘积0的个数来计算答案即可
 **/

func maxTrailingZeros(g [][]int) int {
	m, n := len(g), len(g[0])
	var sum = make([][][2][2]int, m+1)
	for i := 0; i < len(sum); i++ {
		sum[i] = make([][2][2]int, n+1)
	}
	var calc = func(a int) []int {
		var t = a
		cnt2, cnt5 := 0, 0
		for t > 0 && t%2 == 0 {
			t /= 2
			cnt2++
		}
		t = a
		for t > 0 && t%5 == 0 {
			t /= 5
			cnt5++
		}
		return []int{cnt2, cnt5}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			var cur = calc(g[i-1][j-1])
			sum[i][j][0][0] = sum[i-1][j][0][0] + cur[0]
			sum[i][j][0][1] = sum[i-1][j][0][1] + cur[1]
			sum[i][j][1][0] = sum[i][j-1][1][0] + cur[0]
			sum[i][j][1][1] = sum[i][j-1][1][1] + cur[1]
		}
	}
	var res = 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			var u2 = sum[i][j][0][0]
			var u5 = sum[i][j][0][1]
			var l2 = sum[i][j-1][1][0]
			var l5 = sum[i][j-1][1][1]
			var d2 = sum[m][j][0][0] - sum[i-1][j][0][0]
			var d5 = sum[m][j][0][1] - sum[i-1][j][0][1]
			var r2 = sum[i][n][1][0] - sum[i][j][1][0]
			var r5 = sum[i][n][1][1] - sum[i][j][1][1]
			res = max(res, min(u2+l2, u5+l5))
			res = max(res, min(u2+r2, u5+r5))
			res = max(res, min(d2+r2, d5+r5))
			res = max(res, min(d2+l2, d5+l5))
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
