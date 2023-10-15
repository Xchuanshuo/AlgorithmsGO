package test2162

import (
	"fmt"
	"strings"
)

// 1.ts<60, [00:ts]  2.ts>60, [ts/60:ts%60, (ts-60)/60:60 + (ts-60)%60]
func minCostSetTime(startAt int, mc int, pc int, ts int) int {
	var cal = func(s string) int {
		if len(s) > 4 {
			return int(1e9)
		}
		s = strings.TrimLeft(s, "0")
		var last = rune('0' + startAt)
		var cost = 0
		for _, v := range s {
			if v != last {
				cost += mc
			}
			cost += pc
			last = v
		}
		return cost
	}
	if ts < 60 {
		return cal(fmt.Sprintf("%02d%02d", 0, ts))
	}
	l, r := (ts-60)/60, 60+(ts-60)%60
	if r > 99 {
		l += 1
		r -= 60
	}
	var ss = []string{fmt.Sprintf("%02d%02d", ts/60, ts%60),
		fmt.Sprintf("%02d%02d", l, r)}
	return min(cal(ss[0]), cal(ss[1]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
