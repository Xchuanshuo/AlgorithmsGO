package test2151

import "math/bits"

func maximumGood(ss [][]int) int {
	var n = len(ss)
	var l = 1 << n
	var res = 0
	for s := 0; s < l; s++ {
		var valid = true
		for i := 0; i < n; i++ {
			if s&(1<<i) == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if ss[i][j] == 0 && s&(1<<j) != 0 {
					valid = false
				} else if ss[i][j] == 1 && s&(1<<j) == 0 {
					valid = false
				}
				if !valid {
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			res = max(res, bits.OnesCount(uint(s)))
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
