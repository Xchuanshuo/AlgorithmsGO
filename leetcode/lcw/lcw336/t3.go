package lcw336

func beautifulSubarrays(nums []int) int64 {
	var bits = make(map[int]int)
	var mp = make(map[int]int)
	var res int64 = 0
	for _, num := range nums {
		for i := 0; i < 20; i++ {
			if (1<<i)&num > 0 {
				bits[i]++
			}
		}
		var state = 0
		for i := 0; i < 20; i++ {
			var cur = bits[i] % 2
			state |= cur << i
		}
		if v, exist := mp[state]; exist {
			res += int64(v)
		}
		mp[state]++
	}
	return res
}
