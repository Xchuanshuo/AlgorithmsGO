package lcwt115

func getWordsInLongestSubsequence1(n int, words []string, groups []int) []string {
	var dis = func(a, b string) int {
		if len(a) != len(b) {
			return -1
		}
		var cnt = 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				cnt++
			}
		}
		return cnt
	}
	var dp = make([]int, n)
	var parent = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
	max, idx := 0, 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if groups[i] != groups[j] && dp[j]+1 > dp[i] && dis(words[i], words[j]) == 1 {
				dp[i] = dp[j] + 1
				parent[i] = j
			}
			if dp[i] > max {
				max = dp[i]
				idx = i
			}
		}
	}
	var res = make([]string, 0)
	for parent[idx] != -1 {
		res = append(res, words[idx])
		idx = parent[idx]
	}
	for i := 0; i < max/2; i++ {
		res[i], res[max-i-1] = res[max-i-1], res[i]
	}
	return res
}
