package lcw360

import (
	"math"
)

var pos = make(map[int]int)

func init() {
	for i := 0; i < 31; i++ {
		pos[int(math.Pow(2.0, float64(i)))] = i
	}
}

func minOperations(nums []int, target int) int {
	var cnts = make([]int, 33)
	for _, v := range nums {
		cnts[pos[v]]++
	}
	var res = 0
	for i := 0; i < 32; i++ {
		if target&(1<<i) == 0 {
			cnts[i+1] += cnts[i] / 2
			continue
		}
		if cnts[i] > 0 {
			cnts[i]--
			cnts[i+1] += cnts[i] / 2
			continue
		}
		var flag = false
		for j := i + 1; j < 32; j++ {
			if cnts[j] == 0 {
				continue
			}
			flag = true
			res += j - i
			cnts[j]--
			for k := i; k < j; k++ {
				cnts[k]++
			}
			break
		}
		if !flag {
			return -1
		}
		cnts[i+1] += cnts[i] / 2
	}
	return res
}
