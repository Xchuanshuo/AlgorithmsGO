package lcwt113

func minLengthAfterRemovals(nums []int) int {
	var n = len(nums)
	var mp = make(map[int]int)
	var mx = 0
	for _, v := range nums {
		mp[v]++
		mx = max(mx, mp[v])
	}
	if mx*2 > n {
		return mx*2 - n
	}
	return n % 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
