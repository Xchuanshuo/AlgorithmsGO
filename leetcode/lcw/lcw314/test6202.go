package lcw314

func robotWithString(s string) string {
	var n = len(s)
	var f = make([]int, n+1)
	f[n] = 'z'
	for i := n - 1; i >= 0; i-- {
		f[i] = min(f[i+1], int(s[i]))
	}
	var st = make([]int, 0)
	var bytes = make([]byte, 0)
	for i, v := range s {
		st = append(st, int(v))
		for len(st) > 0 && st[len(st)-1] <= f[i+1] {
			bytes = append(bytes, byte(st[len(st)-1]))
			st = st[0 : len(st)-1]
		}
	}
	return string(bytes)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
