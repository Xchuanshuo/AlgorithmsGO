package lcw355

func countPalindromePaths(parent []int, s string) int64 {
	var g = make(map[int][]T)
	for i := 1; i < len(parent); i++ {
		g[parent[i]] = append(g[parent[i]], T{i, int(s[i] - 'a')})
	}
	var res = 0
	var cnt = make(map[int]int)
	cnt[0] = 1
	var dfs func(cur, mask int)
	dfs = func(cur, mask int) {
		for _, nxt := range g[cur] {
			nxtM := mask ^ (1 << nxt.vt)
			res += cnt[nxtM] // nxtM^nxtM
			for i := 0; i < 26; i++ {
				res += cnt[nxtM^(1<<i)] // nxtM ^ x = 1 << i
			}
			cnt[nxtM]++
			dfs(nxt.s, nxtM)
		}
	}
	dfs(0, 0)
	return int64(res)
}

type T struct {
	s, vt int
}
