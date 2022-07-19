package test479

import "math"

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	var max = int64(math.Pow10(n) - 1)
	for i := max; i >= 1; i-- {
		var cur, t int64
		cur, t = i, i
		for t != 0 {
			cur = cur*10 + t%10
			t /= 10
		}
		for j := max; j*j >= cur; j-- {
			if cur%j == 0 {
				return int(cur % 1337)
			}
		}
	}
	return -1
}
