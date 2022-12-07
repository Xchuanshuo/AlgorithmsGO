package test878

/**
 * @Description https://leetcode.cn/problems/nth-magical-number/
 * idea: 二分查找左边界. 确定上边界, n为能被a,b整除的数,
		上边界取r=n*min(a,b), 能保证至少小于r且满足条件的有r/min(a,b) >= n个
		下边界: 二分判定次数即可 小于目标值的数有 能整除a的数 + 能整除b的数 - 能同时整除a和b的数
			 即： mid/a + mid/b - mid/(a,b的最小公倍数)
 **/

func nthMagicalNumber(n int, a int, b int) int {
	l, r := 1, n*min(a, b)
	var M = int(1e9) + 7
	var calc = func(t int) int {
		return t/a + t/b - t/(a/gcd(a, b)*b)
	}
	for l <= r {
		var mid = l + (r-l)/2
		if calc(mid) < n {
			l = mid + 1
		} else {
			if mid == 1 || calc(mid-1) < n {
				return mid % M
			}
			r = mid - 1
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
