package main

import "sort"

/**
 * @Description https://leetcode.cn/contest/biweekly-contest-126/problems/replace-question-marks-in-string-to-minimize-its-value/
 * idea: 贪心
		?处优先用数量较少的字符填充；由于最终分数与位置无关，要使字典序最小，将所有？处的字符排序进行填充即可
 **/

func minimizeStringValue(s string) string {
	var cs = make([]int, 26)
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			continue
		}
		var tv = int(s[i] - 'a')
		cs[tv] += 1
	}
	var ta = make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != '?' {
			continue
		}
		var minV = len(s)
		var tv = 0
		for k, v := range cs {
			if v < minV {
				minV = v
				tv = k
			}
		}
		ta = append(ta, tv)
		cs[tv] += 1
	}
	var bs = []byte(s)
	sort.Ints(ta)
	var k = 0
	for i := 0; i < len(bs); i++ {
		if bs[i] == '?' {
			bs[i] = byte('a' + ta[k])
			k++
		}
	}
	return string(bs)
}
