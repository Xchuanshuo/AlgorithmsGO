package test1665

/**
 * @Description https://leetcode.cn/problems/minimum-initial-energy-to-finish-tasks/
 * idea: 贪心, 先考虑局部，假设有两个任务t1=[a1,m1], t2=[a2,m2]，不同的安排顺序需要不同的初始能量
		 先完成t1,再完成t2，总能量s12 = max(m1, a1+m2); 先完成t2再完成t1，s21 = max(m2, a2+m1)
		 假设 m1-a1 >= m2-a2, 则有 a2+m1 >= a1+m2, 而a为固定消耗的能量，考虑以下几种m的情况即可
		 若m1 > m2, 此时s21=a2+m1，根据假设, 此时满足 s12 <= s21
         若m1 <= m2, 此时s12=a1+m2, 根据假设, a2+m1 >= a1+m2, 则s21=a2+m1, 同样满足s12 <= s21

		 归纳上诉情况可知将mi-ai按降序排列即可，对于全局来说，由于排序后每两个相邻的局部总是满足最优解，全局也为最优解
 **/

import "sort"

func minimumEffort(ts [][]int) int {
	sort.Slice(ts, func(i, j int) bool {
		return ts[i][1]-ts[i][0] > ts[j][1]-ts[j][0]
	})
	cur, res := 0, 0
	for _, t := range ts {
		if cur < t[1] {
			res += t[1] - cur
			cur = t[1]
		}
		cur -= t[0]
	}
	return res
}
