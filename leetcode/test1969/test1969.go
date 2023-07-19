package test1969

var M = int64(1e9) + 7

func minNonZeroProduct(a int) int {
	p := int64(a)
	return int(((1<<p - 1) % M * fastPow(1<<p-2, 1<<(p-1)-1)) % M)
}

func fastPow(a, k int64) int64 {
	if k == 0 {
		return 1
	}
	a %= M
	if k%2 == 0 {
		val := fastPow(a, k/2)
		return (val * val) % M
	} else {
		return a * fastPow(a, k-1) % M
	}
}
