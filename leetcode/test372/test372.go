package test372

const BASE = 1337

func superPow(a int, b []int) int {
	return calc(a, b, len(b) - 1)
}

func calc(a int, b []int, r int) int {
	if r < 0 {
		return 1
	}
	last := b[r]
	cur := fastPow(a, last)
	pre := fastPow(calc(a, b, r-1), 10)
	return cur * pre % BASE
}

func fastPow(a, k int) int {
	if k == 0 {
		return 1
	}
	a %= BASE
	if k%2 == 0 {
		val := fastPow(a, k / 2)
		return (val * val)% BASE
	} else {
		return a * fastPow(a, k - 1) % BASE
	}
}