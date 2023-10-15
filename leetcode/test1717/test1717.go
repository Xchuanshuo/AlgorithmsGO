package test1717

/**
 * @Description https://leetcode.cn/problems/maximum-score-from-removing-substrings
 * idea: 贪心+栈  思路:1.由于存在ab,ba两种不同的子串，所以需要优先抵消分数较大的子串
				   	2.遇到非a、b的字符，前面的子串需要清空，任意一个只包含ab的字符串，要么剩下全为a,要么全为b。
					  否则数量较小的字符都可以与较大的抵消。
		计算分数包含两部分: 1.当前栈顶为较大的分数, 直接出栈累加分数 mxc
						 2.不为较大分数的字符串，则入栈，到非a、b字符清空栈时统计a、b出现的次数，分数为 min(cntA, cntB)*mnc
		为什么第2部分为min(cntA, cntB)*mnc, 因为按照入栈顺序，只要存在较大的分数，都已经进行计算了。最后剩余只能是较小的分数
 **/

func maximumGain(str string, x int, y int) int {
	str += "x"
	var s = make([]byte, 0)
	var res = 0
	mxc, mnc := max(x, y), min(x, y)
	var mp = map[string]int{"ab": x, "ba": y}
	for _, v := range str {
		var b = byte(v)
		if b == 'a' || b == 'b' {
			s = append(s, b)
			if len(s) >= 2 && mp[string(s[len(s)-2:])] == mxc {
				res += mxc
				s = s[0 : len(s)-2]
			}
			continue
		}
		cntA, cntB := 0, 0
		for len(s) > 0 {
			if s[len(s)-1] == 'a' {
				cntA++
			} else {
				cntB++
			}
			s = s[0 : len(s)-1]
		}
		res += min(cntA, cntB) * mnc
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
