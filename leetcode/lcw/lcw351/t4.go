package lcw351

import "sort"

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	var a = make([][]int, 0)
	for i, p := range positions {
		var d = 0 // 0-L, 1-R
		if directions[i] == 'R' {
			d = 1
		}
		a = append(a, []int{i, p, healths[i], d})
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][1] < a[j][1]
	})
	var s = make([][]int, 0)
	for _, cur := range a {
		if len(s) == 0 || !(s[len(s)-1][3] == 1 && cur[3] == 0) {
			s = append(s, cur)
			continue
		}
		// hit
		for len(s) > 0 && s[len(s)-1][3] == 1 && cur[3] == 0 {
			if s[len(s)-1][2] == cur[2] {
				s[len(s)-1][2] = 0
			} else if s[len(s)-1][2] > cur[2] {
				s[len(s)-1][2] -= 1
			} else {
				cur[2] -= 1
				s[len(s)-1] = cur
			}
			if s[len(s)-1][2] == 0 {
				s = s[0 : len(s)-1]
			}
			if !(len(s) > 1 && s[len(s)-2][3] == 1 && s[len(s)-1][3] == 0) {
				break
			}
			cur = s[len(s)-1]
			s = s[0 : len(s)-1]
		}
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i][0] < s[j][0]
	})
	var res = make([]int, 0)
	for _, v := range s {
		res = append(res, v[2])
	}
	return res
}
