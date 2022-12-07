package lcwt91

/**
 * @Description https://leetcode.cn/problems/split-message-based-on-limit
 * idea: 暴力从小到大枚举，将长度为n的字符串切分成i份 记录去除尾部后 剩余的字符串总共的长度cap
		 1.若cap < n, 说明该切分方式不足以容纳整个字符串, 切分份数+1
         2.若cap >= n, 说明找到一种最小的切分方式 模拟构造答案即可
         对于cap, 到达切分份数的边界时, 需要将总容量减去新增的数位
 **/

import "fmt"

func splitMessage(message string, limit int) []string {
	var n = len(message)
	cap, tailLen := 0, 0
	for i := 1; i <= n; i++ {
		if i < 10 {
			tailLen = 5
		} else if i < 100 {
			if i == 10 {
				cap -= 9
			}
			tailLen = 7
		} else if i < 1000 {
			if i == 100 {
				cap -= 99
			}
			tailLen = 9
		} else if i < 10000 {
			if i == 1000 {
				cap -= 999
			}
			tailLen = 11
		}
		cap += limit - tailLen
		if cap < n {
			continue
		}
		var res = make([]string, 0)
		var idx = 0
		for j := 1; j <= i; j++ {
			tail := fmt.Sprintf("<%d/%d>", j, i)
			m := limit - len(tail)
			if idx+m > n {
				res = append(res, message[idx:]+tail)
			} else {
				res = append(res, message[idx:idx+m]+tail)
			}
			idx += m
		}
		return res
	}
	return []string{}
}
