package test2565

/**
 * @Description https://leetcode.cn/problems/subsequence-with-the-minimum-score/
 * idea:  二分 + 前后缀分解
 **/

func minimumScore(s string, t string) int {
	m, n := len(s), len(t)
	var pre = make([]int, n)
	var suf = make([]int, n)
	for i := 0; i < n; i++ {
		pre[i], suf[i] = m, -1
	}
	i, j := 0, 0
	for ; j < n && i < m; i++ {
		if s[i] != t[j] {
			continue
		}
		pre[j] = i
		j++
	}
	i, j = m-1, n-1
	for ; j >= 0 && i >= 0; i-- {
		if s[i] != t[j] {
			continue
		}
		suf[j] = i
		j--
	}
	// fmt.Println(pre, suf)
	// 删除长度为l时能否成为s的子序列
	var check = func(l int) bool {
		for i := 0; i <= n-l; i++ {
			var j = i + l
			pv, sv := -1, m
			if i > 0 {
				pv = pre[i-1]
			}
			if j < n {
				sv = suf[j]
			}
			if pv < sv {
				return true
			}
		}
		return false
	}
	l, r := 0, n
	for l <= n {
		var mid = l + (r-l)/2
		if !check(mid) {
			l = mid + 1
		} else {
			if mid == 0 || !check(mid-1) {
				return mid
			}
			r = mid - 1
		}
	}
	return n
}
