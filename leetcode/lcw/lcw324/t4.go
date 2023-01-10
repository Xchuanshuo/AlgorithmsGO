package lcw324

func cycleLengthQueries(n int, queries [][]int) []int {
	var res = make([]int, 0)
	for _, q := range queries {
		v, w := q[0], q[1]
		var cnt = 0
		for v != w {
			if v < w {
				w /= 2
				cnt++
			} else if v > w {
				v /= 2
				cnt++
			}
		}
		res = append(res, cnt+1)
	}
	return res
}
