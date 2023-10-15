package test2332

/**
 * @Description https://leetcode.cn/problems/the-latest-time-to-catch-a-bus
 * idea: 排序+贪心，最晚到达时间一定是到达时间不连续的乘客中间插空。为了方便计算，可以考虑psgs添加0
		依次考虑每辆车能上的乘客，若前后两个乘客之间有空位，则考虑选择该位置。
		若当前乘客都已经上完，但还有剩余位置，最晚时间则为车到达时间。
 **/

import "sort"

func latestTimeCatchTheBus(buses []int, psgs []int, cap int) int {
	sort.Ints(buses)
	psgs = append(psgs, 0)
	sort.Ints(psgs)
	i, res := 0, 1
	for _, v := range buses {
		var k = cap
		if i == 0 {
			k += 1
		}
		for i < len(psgs) && k > 0 && psgs[i] <= v {
			if i > 0 && psgs[i]-psgs[i-1] > 1 {
				res = psgs[i] - 1
			}
			k--
			i++
		}
		if k > 0 && i > 0 && psgs[i-1] != v {
			res = v
		}
	}
	return res
}
