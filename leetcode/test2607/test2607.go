package test2607

/**
 * @Description https://leetcode.cn/problems/make-k-subarray-sums-equal/
 * idea: a[i]+a[i+1]+...+a[i+k-1] = a[i+1]+a[i+2]+...+a[i+k], 即a[i]=a[i+k]
		 所以要使长度为k的子数组相等，实际是使a[i],a[i+k]..,a[i+nk]的数相等，所以将
		 数字进行分组后排序，计算所有组转换位其中位数所需的次数即可。

		 由于数组还为循环数组，即数组内的数字即有长度为n的周期，又有长度为k的周期，根据裴蜀定理，
		 数组内的数有周期g=gcd(n, k), 直接将元素位置对周期g取模即可快速分组
 **/

import "sort"

func makeSubKSumEqual(arr []int, k int) int64 {
	var n = len(arr)
	var g = gcd(n, k)
	var group = make(map[int][]int)
	for i, v := range arr {
		group[i%g] = append(group[i%g], v)
	}
	var res = int64(0)
	for _, gr := range group {
		sort.Ints(gr)
		var mv = gr[len(gr)/2]
		for _, v := range gr {
			res += int64(abs(v - mv))
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
