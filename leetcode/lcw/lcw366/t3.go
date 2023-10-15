package lcw366

/**
 * @Description https://leetcode.cn/problems/apply-operations-to-make-two-strings-equal
 * idea: 区间dp. 1.不相等的位置数量为奇数时无法使两个字符串相等；若为偶数个，将不等的下标放入pos数组
		        2.考虑区间[i,j], 若长度为2，此时dp[i][j]= min(x, pos[j]-pos[i])
								若长度>2, 此时dp[i][j] = min(dp[i+1][j-1]+x, min(dp[i+2][j]+pos[i+1]-pos[i], dp[i][j-2]+pos[j]-pos[j-1]))
								即要么花费x代价选取首尾下标，要么选取首位相邻的两位下标
 **/
func minOperations(s1 string, s2 string, x int) int {
	var pos = make([]int, 0)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			pos = append(pos, i)
		}
	}
	var n = len(pos)
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return -1
	}
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 2; l <= n; l += 2 {
		for i := 0; i <= n-l; i++ {
			var j = i + l - 1
			if l == 2 {
				dp[i][j] = min(x, pos[j]-pos[i])
			} else {
				dp[i][j] = min(dp[i+1][j-1]+x,
					min(dp[i+2][j]+pos[i+1]-pos[i], dp[i][j-2]+pos[j]-pos[j-1]))
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
