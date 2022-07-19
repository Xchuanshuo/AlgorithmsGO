package test838

func pushDominoes(dominoes string) string {
	var dirs = map[rune]int{'L': -1, 'R': 1}
	var step = make([]int, len(dirs))
	var q = make([]int, 0)
	for idx, ch := range dominoes {
		if ch != '.' {
			step[idx] = 1
			q = append(q, idx)
		}
	}
	var d = []byte(dominoes)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		next := cur + dirs[rune(d[cur])]
		if next < 0 || next >= len(d) {
			continue
		}
		if step[next] == 0 {
			step[next] = step[cur] + 1
			d[next] = d[cur]
			q = append(q, next)
		} else if step[next] == step[cur]+1 {
			d[next] = '.'
		}
	}
	return string(d)
}
