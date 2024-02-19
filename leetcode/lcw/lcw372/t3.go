package lcw372

/**
 * @Description https://leetcode.cn/problems/maximum-xor-product/
 * idea: 贪心.
		1.题目为a，b对同一个数x进行异或后的乘积，讨论对乘积p*q部分的贡献，每位贡献实际都是加上2^i，
		 所以实际 p+q为固定值。要使p*q最大，需要abs(p-q)尽可能小
		2.若a,b的第i位,ba、bb相等，那么该位可以设置为 1 << (ba^1)保证贡献最大
		3.若第i位不相等，则考虑先给较小的数贡献2^i即可
 **/

var M = int64(1e9) + 7

func maximumXorProduct(a int64, b int64, n int) int {
	var p, q int64
	p, q = (a>>n)<<n, (b>>n)<<n
	for i := n - 1; i >= 0; i-- {
		var ba = a >> i & 1
		var bb = b >> i & 1
		if ba == bb {
			p |= 1 << i
			q |= 1 << i
		} else if p > q {
			q |= 1 << i
		} else {
			p |= 1 << i
		}
	}
	p %= M
	q %= M
	return int(p * q % M)
}
