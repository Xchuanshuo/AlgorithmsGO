package test1363

import (
	"sort"
	"strings"
)

func largestMultipleOfThree(digits []int) string {
	var n = len(digits)
	var cnt = make([]int, 3)
	sort.Ints(digits)
	var s = 0
	for _, d := range digits {
		cnt[d%3]++
		s += d
	}
	var rm = func(delCnt, t int) {
		for i := 0; i < n && delCnt > 0; i++ {
			if digits[i]%3 == t {
				delCnt--
				digits[i] = -1
			}
		}
	}
	if s%3 == 1 {
		del, t := 1, 1
		if cnt[1] == 0 {
			del, t = 2, 2
		}
		rm(del, t)
	} else if s%3 == 2 {
		del, t := 1, 2
		if cnt[2] == 0 {
			del, t = 2, 1
		}
		rm(del, t)
	}
	var bs = make([]byte, 0)
	for i := n - 1; i >= 0; i-- {
		if digits[i] == -1 {
			continue
		}
		bs = append(bs, byte(digits[i]+'0'))
	}
	var res = strings.TrimLeft(string(bs), "0")
	if len(bs) > 0 && len(res) == 0 {
		return "0"
	}
	return res
}
