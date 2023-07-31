package test2262

/**
 * @Description https://leetcode.cn/problems/earliest-possible-day-of-full-bloom/
 * idea: 计算新加入当前位置i对前面i-1个字符的贡献即可
 **/

func appealSum(s string) int64 {
	var res int64
	var cur int64
	var last = make([]int, 26)
	for i := 0; i < len(last); i++ {
		last[i] = -1
	}
	for i, v := range s {
		var c = int(v - 'a')
		cur += int64(i - last[c])
		res += cur
		last[c] = i
	}
	return res
}
