package test1392

/**
 * @Description https://leetcode.cn/problems/longest-happy-prefix/
 * idea: 字符串hash
 **/

func longestPrefix(s string) string {
	var h = NewStrHash(s)
	var n = len(s)
	var l = 0
	for i := 1; i < n; i++ {
		if h.Get(1, i) == h.Get(n-i+1, n) {
			l = i
		}
	}
	return s[0:l]
}

type StrHash struct {
	h, p []int
}

func NewStrHash(s string) StrHash {
	var n = len(s)
	h, p := make([]int, n+1), make([]int, n+1)
	p[0] = 1
	for i := 0; i < n; i++ {
		h[i+1] = h[i]*P + int(s[i])
		p[i+1] = p[i] * P
	}
	return StrHash{h, p}
}

const P = 131

func (s StrHash) Get(l, r int) int {
	return s.h[r] - s.h[l-1]*s.p[r-l+1]
}
