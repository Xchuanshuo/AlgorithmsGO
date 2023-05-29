package test1731

var pos = map[int32]int{'a': 0, 'e': 1, 'i': 2, 'o': 3, 'u': 4}

func findTheLongestSubstring(s string) int {
	var state = 0
	var mp = make(map[int]int)
	mp[0] = -1
	var res = 0
	for i, v := range s {
		if _, exist := pos[v]; exist {
			state ^= 1 << pos[v]
		}
		if _, exist := mp[state]; !exist {
			mp[state] = i
		} else {
			res = max(res, i-mp[state])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
