package test866

import (
	"math"
)

/**
 * @Description https://leetcode.cn/problems/prime-palindrome/
 * idea: 从小到大枚举不同长度的回文根即可，先考虑奇数长度的回文数，再考虑偶数
 **/

func primePalindrome(n int) int {
	for L := 1; L <= 5; L++ {
		l, r := int(math.Pow(10, float64(L-1))), int(math.Pow(10, float64(L)))
		// 奇数回文数
		for x := l; x < r; x++ {
			var cur = x
			var last = x / 10
			for last != 0 {
				cur = cur*10 + last%10
				last /= 10
			}
			if cur >= n && isPrim(cur) {
				return cur
			}
		}
		// 偶数回文数
		for x := l; x < r; x++ {
			var cur = x
			var last = x
			for last != 0 {
				cur = cur*10 + last%10
				last /= 10
			}
			if cur >= n && isPrim(cur) {
				return cur
			}
		}
	}
	return 0
}

func isPrim(num int) bool {
	if num == 1 {
		return false
	}
	for j := 2; j <= int(math.Sqrt(float64(num))); j++ {
		if num%j == 0 {
			return false
		}
	}
	return true
}
