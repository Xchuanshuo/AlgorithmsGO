package test2217

import (
	"math"
	"strconv"
)

func kthPalindrome(queries []int, l int) []int64 {
	var res = make([]int64, len(queries))
	var k = (l + 1) / 2
	var base = int(math.Pow10(k - 1))
	for i, q := range queries {
		res[i] = -1
		var t = int64(q - 1 + base)
		var a = t
		if l%2 == 1 {
			a /= 10
		}
		for a != 0 {
			t = t*10 + a%10
			a /= 10
		}
		var lt = len(strconv.Itoa(int(t)))
		if lt == l {
			res[i] = t
		}
	}
	return res
}
