package lcw373

import "strings"

/**
 * @Description https://leetcode.cn/contest/weekly-contest-373/problems/count-beautiful-substrings-ii/
 * idea: 前缀和+哈希表
		1.元音和辅音，转换为+1, -1操作，则问题变为找和为0的子数组个数
		2.对于条件2，设x为任一区间元音字母个数，则需要满足 x^2 % k == 0，考虑如何去掉平方，对k的因子
		 入手，若k为 f * r^2的形式，实际对于x只要有因子r*f即可，等价于 x % (k/r)==0；
		 若k位质数或不包含平方因子，则 x^2 % k == 0等价于 x % k，所以预处理平方即可
		3.根据同余定理，只需要找寻元组为(L/2 % k, s)的子数组即可

 **/

func beautifulSubstrings(s string, k int) int64 {
	var nk = k
	for r := 2; r*r <= k; r++ {
		if k%(r*r) == 0 {
			nk = k / r
		}
	}
	var sum = 0
	var res = int64(0)
	var cnt = make(map[T]int)
	cnt[T{m: nk - 1, s: 0}] = 1
	for i, v := range s {
		if strings.ContainsAny("aeiou", string(v)) {
			sum += 1
		} else {
			sum -= 1
		}
		// 两个条件 1.nk同余 2.sum相等
		var key = T{m: i / 2 % nk, s: sum}
		res += int64(cnt[key])
		cnt[key]++
	}
	return res
}

type T struct {
	m, s int
}