package test1316

// 字符串hash
func distinctEchoSubstrings(s string) int {
	var sh = NewStrHash(s)
	var set = make(map[string]bool)
	for i := 2; i <= len(s); i++ {
		for j := i; j > (i+1)/2; j-- {
			var h1 = sh.Get(j, i)
			var h2 = sh.Get(j-1-(i-j), j-1)
			if h1 == h2 {
				set[s[j-1:i]] = true
			}
		}
	}
	return len(set)
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
