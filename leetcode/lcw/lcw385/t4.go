package lcw385

func countPrefixSuffixPairs(words []string) int64 {
	var res int64
	var cnt = make(map[int]int)
	var mp = make(map[string]int)
	for _, w := range words {
		var hash = NewStrHash(w)
		var n = len(w)
		for i := 1; i <= n-1; i++ {
			var h1 = hash.Get(1, i)
			var h2 = hash.Get(n-i+1, n)
			if h1 == h2 && cnt[h1] > 0 {
				res += int64(cnt[h1])
			}
		}
		if mp[w] > 0 {
			res += int64(mp[w])
		}
		mp[w]++
		cnt[hash.Get(1, n)]++
	}
	return res
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
