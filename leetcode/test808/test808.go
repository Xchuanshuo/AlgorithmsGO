package test808

/**
 * @Description https://leetcode-cn.com/problems/soup-servings/
 * idea:
 *      4个操作每个调整都是25的倍数, 所以可以直接N/25,
 *      再将4个操作分别变为 (-4,-0) (-3,-1) (-2,-2) (-1,-3)
 *      定义dp[i][j]为A剩余i,B剩余j时的概率为多少, 对于剩余量i,j有上述4个操作
 *      所以状态转移方程为
 *      dp[i][j]=0.25*(dp[i-4][j] + dp[i-3][j-1] + dp[i-2][j-2] + dp[i-1][j-3])
 *      初始条件 dp[0][0] A,B都为0时的概率为 A先分完的概率pA=0，A、B同时分完pAB=1 则 0+1/2=0.5
 *              dp[0][1..n]=1, A剩余0，B剩余>0时， pA = 1, pAB=0,  则 1+0/2=1.0
 *              dp[1..n][0]=0, A剩余>0，B位0时, pA=0, pAB=0, 则 0+0/2 = 0
 *      结果 dp[n][n]

 **/

func soupServings(n int) float64 {
	if n >= 5000 {
		return 1.0
	}
	n = (n + 24) / 25
	var dp = make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, n+1)
		dp[i][0] = 0.0
		dp[0][i] = 1.0
	}
	dp[0][0] = 0.5
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = 0.25 * (dp[max(i-4, 0)][j] + dp[max(i-3, 0)][j-1] +
				dp[max(i-2, 0)][max(j-2, 0)] + dp[i-1][max(j-3, 0)])
		}
	}
	return dp[n][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
