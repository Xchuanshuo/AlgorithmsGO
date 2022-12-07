package lcw318

import (
	"math"
	"sort"
)

/**
 * @Description https://leetcode.cn/problems/minimum-total-distance-traveled/
 * idea: 贪心 + 动态规划
		贪心: 机器人一定放到隔自己较近的工厂, 假设放入相邻的工厂 一定需要更多的距离花费
		结论: 对工厂和机器人位置从小到大排序 每个工厂修复的是一段相邻的机器人
       	动态规划: dp[i][j]表示前i个工厂修复前j个机器人所需距离的最小花费
				 枚举当前工厂i修理的机器人个数 [j...k]
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-k] + cost[j...k])
 **/

func minimumTotalDistance(robot []int, fa [][]int) int64 {
	m, n := len(fa), len(robot)
	sort.Ints(robot)
	sort.Slice(fa, func(i, j int) bool {
		return fa[i][0] < fa[j][0]
	})
	var dp = make([][]int64, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int64, n+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt64 / 2
		}
		dp[i][0] = 0
	}
	for i := 1; i <= m; i++ {
		for j := n; j >= 1; j-- {
			var cost int64 = 0
			// 当前工厂不修任何机器人
			dp[i][j] = dp[i-1][j]
			for k := 1; k <= min(fa[i-1][1], j); k++ {
				cost += int64(abs(robot[j-k] - fa[i-1][0]))
				dp[i][j] = min64(dp[i][j], dp[i-1][j-k]+cost)
			}
		}
	}
	return dp[m][n]
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
