package test2122

import "sort"

/**
 * @Description https://leetcode.cn/problems/recover-the-original-array
 * idea: 构造+排序+枚举.
		1.由于nums由原数组+k和-k构成，所以nums至少有一个数对差值d的次数 >= n/2, 且d%2==0
		2.map记录不同差值的次数，过滤后枚举即可。若距离k满足要求，排序后较小
		 的数nums[i]一定能找到一个较大的数nums[i]+2k, 以此构造答案即可
 **/

func recoverArray(nums []int) []int {
	var n = len(nums)
	var mp = make(map[int]int)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			var d = nums[j] - nums[i]
			if d%2 != 0 {
				continue
			}
			mp[d/2]++
		}
	}
	for k, val := range mp {
		if val < n/2 || k == 0 {
			continue
		}
		var t = make(map[int]int)
		for i := 0; i < n; i++ {
			t[nums[i]]++
		}
		var res = make([]int, 0)
		for i := 0; i < n; i++ {
			if t[nums[i]] == 0 {
				continue
			}
			var found = false
			t[nums[i]]--
			if t[nums[i]+2*k] > 0 {
				t[nums[i]+2*k]--
				found = true
			}
			if found {
				res = append(res, nums[i]+k)
			}
		}
		if len(res) == n/2 {
			return res
		}
	}
	return []int{}
}
