package test1808

/**
 * @Description https://leetcode.cn/problems/maximize-number-of-nice-divisors/
 * idea: 质数因子为n个，要好因子数目最大化，实际是求每个质因数分别分配多少个，最后乘积最大
		 即 数为n，拆分成和为n的多段后乘积最大 同 343. 整数拆分
 **/

var M = int(1e9) + 7

func maxNiceDivisors(n int) int {
	var dfs func(n int) int
	dfs = func(n int) int {
		if n <= 4 {
			return n
		}
		if n%3 == 1 {
			return 2 * 2 * pow(3, n/3-1, M) % M
		} else if n%3 == 2 {
			return 2 * pow(3, n/3, M) % M
		}
		return pow(3, n/3, M) % M
	}
	return dfs(n)
}

func pow(a, b, M int) int {
	if b == 0 {
		return 1
	}
	if b&1 == 0 {
		var v = pow(a, b/2, M)
		return v * v % M
	}
	return a * pow(a, b-1, M) % M
}
