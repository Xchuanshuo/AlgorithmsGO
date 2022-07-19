package test1044

var p, h []int64

func longestDupSubstring(s string) string {
	n, base := len(s), int64(1313131)
	p = make([]int64, n + 1)
	h = make([]int64, n + 1)
	p[0] = 1
	for i := 0;i < len(s);i++ {
		p[i+1] = p[i] * base
		h[i+1] = h[i] * base + int64(s[i])
	}
	l, r := 1, len(s)
	for l <= r {
		mid := l + (r-l)/2
		cur := check(s, mid)
		if cur == "" {
			r = mid - 1
		} else {
			if mid == r || check(s, mid + 1) == "" {
				return cur
			}
			l = mid + 1
		}
	}
	return ""
}

func check(s string, l int) string {
	strMap := make(map[int64]int)
	n := len(s)
	for i := 0;i <= n - l;i++ {
		j := i + l - 1
		cur := h[j+1] - h[i]*p[j-i+1]
		strMap[cur]++
		if strMap[cur] > 1 {
			return s[i:j+1]
		}
	}
	return ""
}