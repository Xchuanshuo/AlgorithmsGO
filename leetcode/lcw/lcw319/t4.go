package lcw319

func maxPalindromes(s string, k int) int {
	var n = len(s)
	var isPalind = make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalind[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		isPalind[i][i] = true
	}
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			var j = i + l - 1
			if s[i] == s[j] {
				if l == 2 {
					isPalind[i][j] = true
				} else {
					isPalind[i][j] = isPalind[i+1][j-1]
				}
			}
		}
	}
	var dp = make([]int, n)
	for i := 0; i < n; i++ {
		if i-1 >= 0 {
			dp[i] = dp[i-1]
		}
		for j := i - k + 1; j >= 0; j-- {
			if isPalind[j][i] {
				if j-1 >= 0 {
					dp[i] = max(dp[i], dp[j-1]+1)
				} else {
					dp[i] = max(dp[i], 1)
				}
			}
		}
	}
	return dp[n-1]
}
