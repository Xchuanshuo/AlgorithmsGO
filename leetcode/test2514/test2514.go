package test2514

/**
 * @Description https://leetcode.cn/problems/count-anagrams/
 * idea: 1.统计每个子字符串的排列数 即 n!/x!, 其中x指每个子字符串重复字符出现的次数的阶乘
           由于存在精度损失 所以分子分母分开计算
		 2.每个子字符串的分子分母分别相乘
		 3.使用费马小定理计算最终结果 (a/b)%p = a*b^(p-2)%p 使用快速幂进行计算
 **/

import "strings"

var M = int(1e9) + 7

func countAnagrams(str string) int {
	a, b := 1, 1
	for _, s := range strings.Split(str, " ") {
		var cnt = make(map[int32]int)
		for i, v := range s {
			a = a * (i + 1) % M
			cnt[v]++
			b = b * cnt[v] % M
		}
	}
	return a * fastPow(b, M-2) % M
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

//输入：
//[11,2,19,7,9,27]
//输出：
//31
//预期：
//15
