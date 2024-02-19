package test2949

import "strings"

func beautifulSubstrings(s string, k int) int64 {
	var r = k
	for i := 2; i*i <= k; i++ {
		if k%(i*i) == 0 {
			r = k / i
		}
	}
	var dif = 0
	var res = 0
	var mp = make(map[T]int)
	mp[T{r - 1, 0}] = 1 // r-1与-1同余
	for l, v := range s {
		var x = (l + 1) / 2
		if strings.Contains("aeiou", string(v)) {
			dif += 1
		} else {
			dif -= 1
		}
		var cur = T{dif, x % r}
		res += mp[cur]
		mp[cur]++
	}
	return int64(res)
}

type T struct {
	d, m int
}
