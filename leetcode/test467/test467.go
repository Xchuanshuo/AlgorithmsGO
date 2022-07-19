package test467

func findSubstringInWraproundString(p string) int {
	var dp = make([]int, 26)
	var cnt = 0
	for i, _ := range p {
		if i > 0 && isNext(p[i-1], p[i]) {
			cnt++
		} else {
			cnt = 1
		}
		dp[p[i]-'a'] = max(dp[p[i]-'a'], cnt)
	}
	var res = 0
	for _, v := range dp {
		res += v
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isNext(a, b uint8) bool {
	if a == 'z' && b == 'a' {
		return true
	}
	return a+1 == b
}
