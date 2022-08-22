package test1573

/**
 * @Description
 * idea: 1. sz == 0, C(n-1, 2) = (n-1)!/(2!*(n-3)!) = (n-2)(n-1)/2
 *		 2. sz%3 != 0, 无法分割 0
 *       3. cnt1*cnt2
 **/

func numWays(s string) int {
	var pos = make([]int, 0)
	var M = int(1e9) + 7
	for i, v := range s {
		if v == '1' {
			pos = append(pos, i)
		}
	}
	n, sz := len(s), len(pos)
	if sz == 0 {
		return (n - 1) * (n - 2) / 2 % M
	}
	if sz%3 != 0 {
		return 0
	}
	idx1, idx2 := sz/3, sz/3*2
	cnt1, cnt2 := pos[idx1]-pos[idx1-1], pos[idx2]-pos[idx2-1]
	return cnt1 * cnt2 % M
}
