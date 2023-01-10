package test1691

/**
 * @Description https://leetcode.cn/problems/maximum-height-by-stacking-cuboids
 * idea: 线性dp, 类似LIS,
		前置贪心: 因为下面的长方体严格大于等于上面的 对每个长方体用最长的边作为高度，
		  一定不会比原先摆放方式更加差，所以先将每个长方体的边排序, 再将每个长方体之间进行排序
 **/

import "sort"

func maxHeight(cuboids [][]int) int {
	for _, c := range cuboids {
		sort.Ints(c)
	}
	sort.Slice(cuboids, func(i, j int) bool {
		return cuboids[i][0]+cuboids[i][1]+cuboids[i][2] < cuboids[j][0]+cuboids[j][1]+cuboids[j][2]
	})
	res, n := 0, len(cuboids)
	var dp = make([]int, n)
	for i, c := range cuboids {
		dp[i] = c[2]
		for j := i - 1; j >= 0; j-- {
			if c[0] >= cuboids[j][0] && c[1] >= cuboids[j][1] && c[2] >= cuboids[j][2] {
				dp[i] = max(dp[i], dp[j]+c[2])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
