package lcw356

import (
	"sort"
	"strings"
)

func minimumString(a string, b string, c string) string {
	var calc = func(a, b, c string) string {
		if !strings.Contains(a, b) {
			var idx = 0
			for i := 1; i < len(a); i++ {
				if strings.HasPrefix(b, a[i:]) {
					idx = len(a) - i
					break
				}
			}
			a += b[idx:]
		}
		if !strings.Contains(a, c) {
			var idx = 0
			for i := 1; i < len(a); i++ {
				if strings.HasPrefix(c, a[i:]) {
					idx = len(a) - i
					break
				}
			}
			a += c[idx:]
		}
		return a
	}
	var t []string
	t = append(t, calc(a, b, c))
	t = append(t, calc(a, c, b))
	t = append(t, calc(b, a, c))
	t = append(t, calc(b, c, a))
	t = append(t, calc(c, a, b))
	t = append(t, calc(c, b, a))
	sort.Slice(t, func(i, j int) bool {
		if len(t[i]) == len(t[j]) {
			return t[i] < t[j]
		}
		return len(t[i]) < len(t[j])
	})
	return t[0]
}
