package test375

import "math"

/**
 * @description https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/
 * idea:
 *      dp[i][j]表示在区间[i,j]内, 考虑最差情况猜中数字(即查看完所有数字后才选中) 所需的最小花费是多少
 *      因为要考虑所有情况, 假设当前选中一点k, 需要看k左右两边 哪边完成该游戏需要的
 *      花费较大, 这样才能保证答案在任意一边都可以完成该游戏 而对于整个区间
 *      [i,j]取选取的分割点后最终花费最小的方案 状态转移方程
 *      dp[i][j] = min(dp[i][j], max(dp[i][k-1], dp[k+1][j])+k)
 *      初始条件: 假设只有一个数时, 猜中该数字的花费为0;两个数时, 只需花费较小数值的金币即可
 *
 *      1 2 3
 *      1_2 答案是1: 猜1,0; 猜2,2;   min=1
 *          答案是2: 猜1,1; 猜2,0
 *      2_3 答案是2: 猜2,0; 猜3,3;   min=2
 *          答案是3: 猜2,2; 猜3,0
 *
 *     1_2_3 猜1: 答案是2或3时, max=1 + dp[2][3] = 3
 *           猜2: 答案是1或3时, max=(dp[1][1],dp[3][3])+2 = 2
 *           猜3: 答案是1或2时, max=3 + dp[1][2] = 4
 *        三种情况取最小值 即dp[1][3]=2
 **/

func getMoneyAmount(n int) int {
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt
		}
		dp[i][i] = 0
	}
	for l := 2; l <= n+1; l++ {
		for i := 0; i <= n-l+1; i++ {
			j := i + l - 1
			if l == 2 {
				dp[i][j] = i
				continue
			}
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], max(dp[i][k-1], dp[k+1][j])+k)
			}
		}
	}
	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
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
