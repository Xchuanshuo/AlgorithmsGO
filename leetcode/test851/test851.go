package test851

func loudAndRich(richer [][]int, quiet []int) []int {
	degree := make([]int, len(quiet))
	rG := make(map[int][]int)
	for _, r := range richer {
		rG[r[0]] = append(rG[r[0]], r[1])
		degree[r[1]]++
	}
	q := make([]int, 0, len(quiet))
	ans := make([]int, len(quiet))
	for idx, v := range degree {
		ans[idx] = idx
		if v == 0 {
			q = append(q, idx)
		}
	}
	for len(q) > 0 {
		y := q[0]
		q = q[1:]
		for _, x := range rG[y] {
			if quiet[ans[y]] < quiet[ans[x]] {
				ans[x] = ans[y]
			}
			degree[x]--
			if degree[x] == 0 {
				q = append(q, x)
			}
		}
	}
	return ans
}