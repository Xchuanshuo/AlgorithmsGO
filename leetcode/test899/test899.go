package test899

import "sort"

/**
 * @Description https://leetcode.cn/problems/orderly-queue/
 *  	k>1 时可以交换任意相邻顺序，返回字典序列
 *	    k=1 时，模拟构造序列即可
 **/

func orderlyQueue(s string, k int) string {
	if k == 1 {
		var res = s
		for i := 1; i < len(s); i++ {
			var cur = s[i:] + s[0:i]
			if cur < res {
				res = cur
			}
		}
		return res
	}
	var bytes = []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}
