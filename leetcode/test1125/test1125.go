package test1125

func smallestSufficientTeam(skills []string, people [][]string) []int {
	m, n := len(people), len(skills)
	var mp = make(map[string]int)
	for i, s := range skills {
		mp[s] = i
	}
	var pv = make([]int, m)
	for i, p := range people {
		var v = 0
		for _, s := range p {
			v |= 1 << mp[s]
		}
		pv[i] = v
	}
	var l = 1 << n
	var dp = make([][]int, l)
	dp[0] = make([]int, 0)
	for i := 0; i < l; i++ {
		if dp[i] == nil {
			continue
		}
		for j, v := range pv {
			if i&v == v {
				continue
			}
			var nxt = i | v
			if dp[nxt] == nil || len(dp[nxt]) > len(dp[i])+1 {
				var dst = make([]int, len(dp[i]))
				copy(dst, dp[i])
				dp[nxt] = append(dst, j)
			}
		}
	}
	return dp[l-1]
}
