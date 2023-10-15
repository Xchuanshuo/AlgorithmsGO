package test1589

import "sort"

var M = int(1e9) + 7

// 1.差分数组更新所有区间的查询次数 2.所有位置更新次数排序 3.依次放入最大值
func maxSumRangeQuery(nums []int, requests [][]int) int {
	sort.Ints(nums)
	var n = len(nums)
	var dif = make([]int, n+1)
	for _, r := range requests {
		dif[r[0]] += 1
		dif[r[1]+1] -= 1
	}
	for i := 1; i < len(dif); i++ {
		dif[i] += dif[i-1]
	}
	sort.Ints(dif)
	res, r := 0, n-1
	for i := n; i >= 0; i-- {
		if dif[i] == 0 {
			break
		}
		res = (res + dif[i]*nums[r]) % M
		r--
	}
	return res
}
