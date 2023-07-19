package lcwt108

import (
	"sort"
)

var dirs = [][]int{{-1, 1}, {-1, -1}, {1, 1}, {1, -1}}

type t struct {
	x1, y1, x2, y2 int
}

func countBlackBlocks(m int, n int, coordinates [][]int) []int64 {
	var total = int64(m-1) * int64(n-1)
	var idMap = make(map[t]int)
	for _, p := range coordinates {
		x, y := p[0], p[1]
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx >= m || nx < 0 || ny < 0 || ny >= n {
				continue
			}
			var ps = [][]int{{x, y}, {nx, ny}}
			sort.Slice(ps, func(i, j int) bool {
				return ps[i][0] < ps[j][0]
			})
			if d[0] != d[1] {
				ps = [][]int{{ps[0][0], ps[0][1] - 1}, {ps[1][0], ps[1][1] + 1}}
			}
			idMap[t{ps[0][0], ps[0][1], ps[1][0], ps[1][1]}]++
		}
	}
	var res = make([]int64, 5)
	for _, v := range idMap {
		res[v]++
	}
	res[0] = total - int64(len(idMap))
	return res
}
