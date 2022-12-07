package lcwt92

/**
 * @Description https://leetcode.cn/problems/count-palindromic-subsequences
 * idea:
        1.统计所有长度为5的序列 => 统计以当前字符为中心，左右各取两个字符能构成回文的方案数
		2.因为字符范围为0-9, 所以两两组合能构成回文的方案数为100
        3.左右两遍扫描，记录00-99每个序列出现的次数，使用乘法原理即可得出答案
 **/

func countPalindromes(s string) int {
	n, M := len(s), int(1e9)+7
	var suf = make([][]int, n)
	var cnt = make([]int, 10)
	var t = make([]int, 100)
	for i := n - 1; i >= 0; i-- {
		var c = int(s[i] - '0')
		for k := 0; k < 10; k++ {
			if cnt[k] != 0 {
				t[c*10+k] += cnt[k]
			}
		}
		var dst = make([]int, 100)
		copy(dst, t)
		suf[i] = dst
		cnt[c]++
	}
	cnt = make([]int, 10)
	t = make([]int, 100)
	var res = 0
	for i := 0; i < n-1; i++ {
		var c = int(s[i] - '0')
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				res = (res + t[j*10+k]*suf[i+1][j+10*k]) % M
			}
		}
		for k := 0; k < 10; k++ {
			if cnt[k] != 0 {
				t[k*10+c] += cnt[k]
			}
		}
		cnt[c]++
	}
	return res
}
