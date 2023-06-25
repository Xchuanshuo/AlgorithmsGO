package lcwt106

import "sort"

var Mod = int(1e9) + 7
func sumDistance(nums []int, s string, d int) int {
	var n = len(nums)
	for i := 0;i < n;i++ {
		if s[i] == 'L' {
			nums[i] -= d
		} else {
			nums[i] += d
		}
	}
	sort.Ints(nums)
	var res = 0
	var sum = 0
	for i := 0;i < n;i++ {
		res = (res + i*nums[i] - sum)%Mod
		sum = (sum+nums[i])%Mod
	}
	return res
}
