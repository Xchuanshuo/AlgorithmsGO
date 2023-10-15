package test2311

import (
	"strconv"
)

// 贪心, 1.尽可能的加入更多前缀0   2.对于s的len(k)位后缀，分情况讨论，
//		若后len(k)位 <= k, 则总长度为前面0的个数 + len(k)
//		若后len(k)位 > k, 则总长度为前面0的个数 + len(k) - 1, -1表示去掉后len(k)位的任意一位

func longestSubsequence(s string, k int) int {
	var sk = toBinStr(k)
	if len(sk) > len(s) {
		return len(s)
	}
	sv, _ := strconv.ParseInt(s[len(s)-len(sk):], 2, 64)
	var res = cnt0(s[0:len(s)-len(sk)]) + len(sk)
	if int(sv) <= k {
		return res
	}
	return res - 1
}

func toBinStr(a int) string {
	var bs []byte
	for a != 0 {
		bs = append(bs, byte('0'+a%2))
		a /= 2
	}
	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(bs)-i-1] = bs[len(bs)-i-1], bs[i]
	}
	return string(bs)
}

func cnt0(s string) int {
	var cnt = 0
	for _, v := range s {
		if v == '0' {
			cnt++
		}
	}
	return cnt
}
