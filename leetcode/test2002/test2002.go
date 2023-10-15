package test2002

/**
 * @Description https://leetcode.cn/problems/maximum-product-of-the-length-of-two-palindromic-subsequences
 * idea: 状态压缩+dp. 枚举所有状态，分别求当前状态和取反后状态的最大回文子序列长度，计算答案即可
 **/

func maxProduct(str string) int {
	var n = len(str)
	var l = 1 << n
	var dp = make([]int, l)
	for i := 1; i < len(dp); i++ {
		dp[i] = 1
	}
	var res = 0
	for s := 1; s < l; s++ {
		var bs = make([]byte, 0)
		for i := 0; i < n; i++ {
			if s&(1<<i) != 0 {
				bs = append(bs, str[i])
			}
		}
		dp[s] = cal(string(bs))
		res = max(res, dp[s]*dp[s^(l-1)])
	}
	return res
}

var cal = func(s string) int {
	var n = len(s)
	var dp = make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			var j = i + l - 1
			if l == 1 {
				dp[i][j] = 1
				continue
			}
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
