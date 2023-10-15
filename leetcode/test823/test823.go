package test823

/**
 * @Description https://leetcode.cn/problems/binary-trees-with-factors/
 * idea: 排序+动态规划 dp[v]表示以v为乘积的子树个数，在arr内枚举因子l,r 则dp[v] += dp[l]*dp[r],
					若因子l或r不存在 乘积为0，对答案无影响
 **/
import "sort"

var M = int(1e9) + 7

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	var res = 0
	var dp = make(map[int]int)
	for i, v := range arr {
		dp[v] = 1
		for j := 0; j < i; j++ {
			if v%arr[j] != 0 {
				continue
			}
			l, r := arr[j], v/arr[j]
			dp[v] = (dp[v] + dp[l]*dp[r]) % M
		}
		res = (res + dp[v]) % M
	}
	return res
}
