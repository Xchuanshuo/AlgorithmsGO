package test1201

/**
 * @Description https://leetcode.cn/problems/ugly-number-iii/
 * idea: n范围较大，考虑二分。计算数字1..t能构造多少个丑数，即范围内能
		被a,b,c整除的不同数字有多少个，使用容斥原理。分为几部分
		1.整除a的数字，整除b的数字，整除c的数字。即t/a + t/b +t/c
		2.1中同时能被a,b,c两个数以上整除的数重复计算了，需要剔除。即 t/lcm(a,b) + t/lcm(a,c) + t/lcm(b,c)
		3.2的剔除方式会导致同时被a,b,c整除的数(1-重复3次，2-减去3次)都被剔除了，需要加回。即t/lcm(lcm(a, b), c)
 **/

func nthUglyNumber(n int, a int, b int, c int) int {
	var mx = int(2e9)
	var valid = func(t int) bool {
		var cnt = t/a + t/b + t/c - (t/lcm(a, b) + t/lcm(a, c) + t/lcm(b, c)) + t/lcm(lcm(a, b), c)
		return cnt >= n
	}
	l, r := 1, mx
	for l <= r {
		var mid = l + (r-l)/2
		if !valid(mid) {
			l = mid + 1
		} else {
			if mid == 1 || !valid(mid-1) {
				return mid
			}
			r = mid - 1
		}
	}
	return 0
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
