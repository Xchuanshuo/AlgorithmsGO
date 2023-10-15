package test2327

/**
 * @Description https://leetcode.cn/problems/number-of-people-aware-of-a-secret/
 * idea:  dp[j]表示第j天新知道秘密的人数, 枚举当前天数i, 计算天数j新知道秘密的人数，当前新知道秘密
         的人在i+delay天后会分享秘密 在i+forget会忘记秘密 所以 dp[j] += dp[i]，j的范围为[i+delay, i+forget-1]
		答案为: sum(dp[n-forget+1], dp[n]), 即到第n天为止还没忘记秘密的人生
 **/

var M = int(1e9) + 7

func peopleAwareOfSecret(n int, delay int, forget int) int {
	var dp = make([]int, n+1+delay+forget)
	dp[1] = 1
	for i := 1; i <= n; i++ {
		for j := i + delay; j < i+forget; j++ {
			dp[j] = (dp[j] + dp[i]) % M
		}
	}
	var res = 0
	for i := n - forget + 1; i <= n; i++ {
		res = (res + dp[i]) % M
	}
	return res
}
