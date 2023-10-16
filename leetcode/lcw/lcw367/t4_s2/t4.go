package t4_s2

/**
 * @Description https://leetcode.cn/problems/construct-product-matrix
 * idea: 乘法逆元。根据题意，很自然的想到思路，当前位置的答案为 全部数的(乘积/当前数)%12345
		 大整数除法取模不满足同余定理，考虑转换成乘法，求逆元。由于12345非质数，无法使用
		 费马小定理计算逆元，所以使用扩展欧几里得算法来求解。扩展欧几里得需要满足gcd(a,b)=1,
		 且12345=3*5*823, 所以需要将矩阵中的数去除掉因子3、5、823，才能使该数与12345互质。

		 具体步骤如下
		 1.对每个数的3、5、823因子进行计数
		 2.将乘积分为两部分，一部分为所有不包含3、5、823因子的乘积；另一部分为所有这三个因子的乘积。
		  前一部分将除当前数求模转换为乘以逆元求模，第二部分去除当前位置的因子数后，使用快速幂，两者相乘即为最后的答案。
 **/

var M = 12345

func constructProductMatrix(g [][]int) [][]int {
	m, n := len(g), len(g[0])
	// 分别记录因子为3,5,823的出现次数
	s3, s5, s823 := 0, 0, 0
	// 返回去除因子后的数
	var div = func(t int) (int, int, int, int) {
		c3, c5, c823 := 0, 0, 0
		for t%3 == 0 {
			c3++
			t /= 3
		}
		for t%5 == 0 {
			c5++
			t /= 5
		}
		for t%823 == 0 {
			c823++
			t /= 823
		}
		return t, c3, c5, c823
	}
	var total = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			t, c3, c5, c823 := div(g[i][j])
			total = (total * t) % M
			s3 += c3
			s5 += c5
			s823 += c823
		}
	}
	var res = make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			t, c3, c5, c823 := div(g[i][j])
			var x = inv(t, M)                            // x为t的逆元
			if s3-c3 > 0 && s5-c5 > 0 && s823-c823 > 0 { // 存在12345的三个质因子
				res[i][j] = 0
			} else {
				res[i][j] = total * x % M *
					fastPow(3, s3-c3) % M *
					fastPow(5, s5-c5) % M *
					fastPow(823, s823-c823) % M
			}
		}
	}
	return res
}

func fastPow(a, k int) int {
	if k == 0 {
		return 1
	}
	a %= M
	if k%2 == 0 {
		val := fastPow(a, k/2) % M
		return val * val % M
	}
	return a * fastPow(a, k-1) % M
}

// 扩展欧几里得算法, 返回gcd, x, y
func extGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	d, x, y := extGCD(b, a%b)
	return d, y, x - (a/b)*y
}

// 求逆元 a, p互质的情况下才可以求解
func inv(a, p int) int {
	d, x, _ := extGCD(a, p)
	if d != 1 {
		panic("逆元不存在")
	}
	return (x + p) % p // 防止x为负数
}
