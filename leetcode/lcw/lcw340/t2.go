package lcw340


func distance(nums []int) []int64 {
	var mp = make(map[int][]int)
	for i, v := range nums {
		mp[v] = append(mp[v], i)
	}
	var res = make([]int64, len(nums))
	for _, ls := range mp {
		var n = len(ls)
		var s = make([]int, n+1)
		for i, v := range ls {
			s[i+1] = s[i] + v
		}
		for i, v := range ls {
			sl, sr := s[i], s[n]-s[i+1]
			res[v] += int64(v*i-sl) + int64(sr - v*(n-i-1))
		}
	}
	return res
}
