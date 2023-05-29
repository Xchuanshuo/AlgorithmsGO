package test1092

func shortestCommonSupersequence(s1 string, s2 string) string {
	m, n := len(s1), len(s2)
	var dp = make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	i, j := m, n
	var bs = make([]byte, 0)
	for i > 0 || j > 0 {
		if i <= 0 {
			bs = append(bs, s2[j-1])
			j--
		} else if j <= 0 {
			bs = append(bs, s1[i-1])
			i--
		} else {
			if s1[i-1] == s2[j-1] {
				bs = append(bs, s1[i-1])
				i--
				j--
			} else if dp[i-1][j] == dp[i][j] {
				bs = append(bs, s1[i-1])
				i--
			} else {
				bs = append(bs, s2[j-1])
				j--
			}
		}
	}
	var l = len(bs)
	for i := 0; i < l/2; i++ {
		bs[i], bs[l-i-1] = bs[l-i-1], bs[i]
	}
	return string(bs)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
