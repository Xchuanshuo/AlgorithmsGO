package lcw359

func minimumSum(n int, k int) int {
	var mp = make(map[int]bool)
	var res = 0
	var cnt = 0
	for i := 1; i <= 2*n; i++ {
		if cnt >= n {
			break
		}
		if mp[k-i] {
			continue
		}
		res += i
		cnt++
		mp[i] = true
	}
	return res
}
