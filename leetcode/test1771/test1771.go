package test1771

/**
 * @Description https://leetcode.cn/problems/maximize-palindrome-length-from-subsequences
 * idea: 区间dp,同最长回文子序列, 关键点: 找出满足左右边界i,j分别出现在两个字符串位置才能计算答案
 **/

func longestPalindrome(word1 string, word2 string) int {
	var s = word1 + word2
	n, m := len(word1), len(word2)
	var total = n + m
	var dp = make([][]int, total)
	for i := 0; i < total; i++ {
		dp[i] = make([]int, total)
	}
	var res = 0
	for l := 1; l <= total; l++ {
		for i := 0; i <= total-l; i++ {
			var j = i + l - 1
			if l == 1 {
				dp[i][j] = 1
			} else if l == 2 {
				if s[i] == s[j] {
					dp[i][j] = 2
				} else {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1] + 2
				} else {
					dp[i][j] = max(dp[i+1][j], dp[i][j-1])
				}
			}
			if s[i] == s[j] && (i < n && j >= n) {
				res = max(res, dp[i][j])
			}
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
