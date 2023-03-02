package lcw333

import "math"

var M = int(1e9) + 7
var mp = make(map[int][]int)

func init() {
	for i := 2;i <= 30;i++ {
		mp[i] = f(i)
	}
}

func squareFreeSubsets(nums []int) int {
	var cnt = 0
	var n = len(nums)


	return fastPow(2, n-cnt) - 1
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

var f = func(n int) []int {
	var facs = make([]int, 0)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
				facs = append(facs, i)
			}
		}
	}
	if n >= 1 {
		facs = append(facs, n)
	}
	return facs
}
