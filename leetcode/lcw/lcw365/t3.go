package lcw365

import "math"

func minSizeSubarray(nums []int, tmp int) int {
	var target = int64(tmp)
	var n = len(nums)
	var sum = int64(0)
	for _, v := range nums {
		sum += int64(v)
	}
	tl, target := target/sum*int64(n), target%sum
	if target == 0 {
		return int(tl)
	}
	nums = append(nums, nums...)
	var mp = make(map[int64]int)
	mp[0] = -1
	var res = math.MaxInt32
	for i, v := range nums {
		sum += int64(v)
		if p, exist := mp[sum-target]; exist {
			res = min(res, i+int(tl)-p)
		}
		mp[sum] = i
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
