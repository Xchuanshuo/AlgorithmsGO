package lcwt124

/**
 * @Description https://leetcode.cn/contest/biweekly-contest-124/problems/maximize-consecutive-elements-in-an-array-after-modification/
 * idea: 排序 + dp, 计算每个位置变或不变的元素序列长度
 **/

import "sort"

func maxSelectedElements(nums []int) int {
	sort.Ints(nums)
	var mp = make(map[int]int)
	var res = 1
	for i := len(nums) - 1; i >= 0; i-- {
		var v = nums[i]
		mp[v] = max(1+mp[v+1], mp[v])
		mp[v+1] = max(mp[v+2]+1, mp[v+1])
		res = max(res, mp[v], mp[v+1])
	}
	return res
}
