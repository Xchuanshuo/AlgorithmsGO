package lcw351

import "math/bits"

func makeTheIntegerZero(num1 int, num2 int) int {
	for k := 0; num1-k*num2 > 0; k++ {
		var cur = num1 - k*num2
		if cur >= k && bits.OnesCount(uint(cur)) <= k {
			return k
		}
	}
	return -1
}
