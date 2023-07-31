package test1888

/**
 * @Description https://leetcode.cn/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/
 * idea: 1.操作1，将首字符删除后移到末尾，问题可以转换为原字符串拼接原字符串后，求长度为n区间的最小翻转次数
		 2.答案交替字符串可能为1或0结尾，所以需求分情况讨论，再根据n的奇偶性，可以得出当前字符前n+1位置的字符是什么，
		   操作次数减去这一部分即可。
 **/

func minFlips(s string) int {
	var n = len(s)
	s += s
	var dp = make([][2]int, len(s)+1)
	var res = n
	for i := 1; i <= len(s); i++ {
		if s[i-1] == '0' {
			dp[i][0] = dp[i-1][1]
			dp[i][1] = dp[i-1][0] + 1
		} else {
			dp[i][0] = dp[i-1][1] + 1
			dp[i][1] = dp[i-1][0]
		}
		if i >= n {
			res = min(res, min(dp[i][0]-dp[i-n][n%2], dp[i][1]-dp[i-n][n%2^1]))
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
