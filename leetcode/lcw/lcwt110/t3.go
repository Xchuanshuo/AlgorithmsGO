package lcwt110

func minimumSeconds(nums []int) int {
	nums = append(nums, nums...)
	var n = len(nums)
	var pos = make(map[int]int)
	var dis = make(map[int]int)
	for i := 0; i < n/2; i++ {
		pos[nums[i]] = i
	}
	for i, v := range nums {
		if pos[v] == -1 {
			dis[v] = i
		} else {
			dis[v] = max(dis[v], (i-pos[v])/2)
		}
		pos[v] = i
	}
	var res = n
	for _, v := range dis {
		res = min(res, v)
	}
	return res
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
