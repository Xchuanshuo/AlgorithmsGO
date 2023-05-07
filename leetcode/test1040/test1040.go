package test1040

import "sort"

// 1.最大 区间空位 - (总长度去除两个端点)
// 2.最小 长度为n的窗口剩余最少空位 -> n-窗口最大石子数
func numMovesStonesII(s []int) []int {
	sort.Ints(s)
	var n = len(s)
	s1, s2 := s[n-1]-s[1]-(n-2), s[n-2]-s[0]-(n-2)
	mx := max(s1, s2)
	if s1 == 0 || s2 == 0 {
		return []int{min(2, mx), mx}
	}
	l, cnt := 0, 0
	for r := 0; r < n; r++ {
		for s[r]-s[l]+1 > n {
			l++
		}
		cnt = max(cnt, r-l+1)
	}
	return []int{n - cnt, mx}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
