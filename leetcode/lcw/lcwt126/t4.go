package main

import (
	"fmt"
	"maps"
)

/**
 * @Description https://leetcode.cn/contest/biweekly-contest-126/problems/find-the-sum-of-the-power-of-all-subsequences/
 * idea: dp + 快速幂 + 计数
		1.计算所有子序列和为k的能量 => 选定和为k长度为l的序列，其它数任选，即 cnt * 2^(n-l)
		2.由于不同长度的序列对计算答案有影响，所以需要计数的key为(sum,l)
 **/

var M = 1000000007

func sumOfPower(nums []int, k int) int {
	var preSums = make(map[T]int)
	preSums[T{0, 0}] = 1
	for _, num := range nums {
		var newSums = make(map[T]int)
		maps.Copy(newSums, preSums)
		for pre, cnt := range preSums {
			if pre.s+num > k { // 大于k无需计算
				continue
			}
			var nxt = T{pre.s + num, pre.l + 1}
			newSums[nxt] = (newSums[nxt] + cnt) % M
		}
		preSums = newSums
	}
	var res = 0
	var n = len(nums)
	for t, cnt := range preSums {
		if t.s == k {
			res = (res + cnt*fastPow(2, n-t.l)) % M
		}
	}
	return res
}

type T struct {
	s, l int
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
func main() {
	nums1 := []int{1, 2, 3}
	k1 := 3
	fmt.Printf("能量 for nums1: %d\n", sumOfPower(nums1, k1))

	nums2 := []int{2, 3, 3}
	k2 := 5
	fmt.Printf("能量 for nums2: %d\n", sumOfPower(nums2, k2))
}
