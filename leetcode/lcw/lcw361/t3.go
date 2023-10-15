package lcw361

// 取模 + 同余
func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	var n = len(nums)
	var sum = make([]int, n+1)
	var res = int64(0)
	for i, v := range nums {
		var c = 0
		if v%modulo == k {
			c = 1
		}
		sum[i+1] = sum[i] + c
	}
	var mp = make(map[int]int)
	for _, v := range sum {
		var cur = v % modulo
		// s1%t, (s2%t-k)%t => (s2-s1)%k == t
		if v, exist := mp[(cur-k+modulo)%modulo]; exist {
			res += int64(v)
		}
		mp[cur]++
	}
	return res
}
