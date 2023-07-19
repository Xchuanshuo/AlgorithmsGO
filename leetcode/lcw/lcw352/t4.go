package main

func sumImbalanceNumbers(nums []int) int {
	var n = len(nums)
	var res = 0
	for i := 0; i < n; i++ {
		var cnt = 0
		var mp = make(map[int]int)
		mp[nums[i]] = 1
		for j := i + 1; j < n; j++ {
			if mp[nums[j]] == 0 {
				mp[nums[j]] = 1
				cnt += 1 - mp[nums[j]-1] - mp[nums[j]+1]
			}
			res += cnt
		}
	}
	return res
}
