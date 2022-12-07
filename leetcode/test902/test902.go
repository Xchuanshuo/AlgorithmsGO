package test902

/**
 * @Description https://leetcode-cn.com/problems/numbers-at-most-n-given-digit-set/
 * idea:
 *      数位dp dp[i]表示对于数字前i位(从低位到高位)的数字组合个数
 *      对于n位的数字, 前n-1位是可以使用digits中的数任意组合的, 组合个数为lenD^(n-1)
 *      以 digits={1,2,3,4} n = 2345 为例 n总共有4位            2 3 4 5
 *                                                        位: 4 3 2 1
 *    ------------------------------------------------------------------------------
 *      总方案数包括两部分 1.前面i位置0 剩余位任意取 2.(不置0)当前位改变后的(digits的数 小于 或 等于当前位两种情况)方案汇总
 *
 *      先看第1位, digits中有4个数小于5的, 则前1位有[4]种方案
 *      再看第2位, digits中有3个小于4的, 取这三个数后 则前1位可以任意取 方案数为 [3*4^1=12]
 *               digits中有1数个等于4, 取这1个数后, 则方案数实际为dp[1]得到的方案数 两种情况累加即可
 *      接下来看第3位, 有2个数小于3, 当前位取这两个数后, 前2位可以任意组合 即 [2*4^2=32]
 *                   有1个数等于3, 当前位取3后, 方案数为dp[2] 两种情况累加
 *      最后看第4位, 有1个数小于2, 则当前位取这一个数后, 方案为 [1*4^3=64]
 *                 有1个数等于2, 则当前位取2后, 方案为dp[3] 两种情况累加
 *      则dp[1] = 4, dp[2]=12 + 4 = 16, dp[3]= 32 + 16 = 48
 *      前4位方案数为 dp[4] = dp[3] + 64 = 112
 *      还需要考虑[当前位不动] 前n-1位任意取的情况  4^1 + 4^2 + 4^3 = 84
 *      则方案总数为 dp[4] + 84 = 196
 **/

import "math"

func atMostNGivenDigitSet(digits []string, n int) int {
	var t = n
	var f = make([]int, 0)
	for t != 0 {
		f = append(f, t%10)
		t /= 10
	}
	var m = len(f)
	var dp = make([]int, m+1)
	dp[0] = 1
	for i := 1; i <= m; i++ {
		var cur = f[i-1]
		for _, d := range digits {
			var v = int(d[0] - '0')
			if v < cur {
				dp[i] += int(math.Pow(float64(len(digits)), float64(i-1)))
			} else if v == cur {
				dp[i] += dp[i-1]
			}
		}
	}
	for i := 1; i < m; i++ {
		dp[m] += int(math.Pow(float64(len(digits)), float64(i)))
	}
	return dp[m]
}
