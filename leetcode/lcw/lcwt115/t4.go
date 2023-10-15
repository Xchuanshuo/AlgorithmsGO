package lcwt115

/**
 * @Description https://leetcode.cn/contest/biweekly-contest-115/problems/count-of-sub-multisets-with-bounded-sum/
 * idea: 分组背包 + 前缀和优化。
		   1.根据数据范围和等差数列求和公式，可以得出nums值域的范围为[0,sqrt(sum)]
		   2.分组背包问题。dp[i][j]表示前i个数和为j的方案数。由于对于一个数x能取任意次，
			直接枚举所有选法会超时，所以考虑对计算结果进行优化。
		   3.对转移方程入手，当前选取的数为x时的方案数为如下
			1).dp[i][j] = dp[i-1][j] + dp[i-1][j-x] + dp[i-1][j-2*x] +...+ dp[i-1][j-cnt[x]*x]
		 而 2).dp[i][j-x] = dp[i-1][j-x] + dp[i-1][j-2*x] + ... + dp[i-1][j-(cnt[x]+1)*x]
		  可以观察到1)比2)式多了一个dp[i-1][j], 少了一个dp[i-1][j-(cnt[x]+1)*x], 所以对于当前方案数
			 dp[i][j] = dp[i-1][j] + dp[i][j-x] - dp[i-1][j-(cnt[x]+1)*x]

		 初始条件 dp[i][0]表示前i个元素和为0的方案数，空集 + 选[1,n]个0的集合，即 cnt[0]+1
		 答案为 sum(dp[n][l]..dp[n][r])
 **/

var M = int(1e9) + 7

func countSubMultisets(nums []int, l int, r int) int {
	var cnt = make(map[int]int)
	for _, x := range nums {
		cnt[x]++
	}
	var n = len(cnt)
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, r+1)
	}
	dp[0][0] = cnt[0] + 1 // 初始条件 空集 + 0的个数
	var i = 1
	for x, c := range cnt {
		for j := 0; j <= r; j++ {
			dp[i][j] = dp[i-1][j]
			if j-x < 0 {
				continue
			}
			var sub = 0
			var t = j - (c+1)*x
			if t >= 0 {
				sub = dp[i-1][t]
			}
			dp[i][j] = (dp[i][j] + dp[i][j-x] - sub + M) % M
		}
		i++
	}
	var res = 0
	for i := l; i <= r; i++ {
		res = (res + dp[n][i]) % M
	}
	return res
}
