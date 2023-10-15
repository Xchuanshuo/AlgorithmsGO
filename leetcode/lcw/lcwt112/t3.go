package lcwt112

func maxSum(nums []int, m int, k int) int64 {
	var set = make(map[int]int)
	var sum = int64(0)
	for i := 0; i < k; i++ {
		sum += int64(nums[i])
		set[nums[i]]++
	}
	var res = int64(0)
	if len(set) >= m {
		res = sum
	}
	for i := k; i < len(nums); i++ {
		var lv = nums[i-k]
		sum -= int64(lv)
		set[lv]--
		if set[lv] == 0 {
			delete(set, lv)
		}
		var rv = nums[i]
		set[rv]++
		sum += int64(rv)
		if len(set) >= m {
			res = max(res, sum)
		}
	}
	return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
