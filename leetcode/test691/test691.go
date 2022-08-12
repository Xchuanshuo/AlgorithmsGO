package test691

func minStickers(stickers []string, t string) int {
	n, m := len(stickers), len(t)
	var INF = int(1e9)
	var cnts = make([][]int, n)
	var g = make([][]int, 26)
	for i := 0; i < n; i++ {
		cnts[i] = make([]int, 26)
		for _, v := range stickers[i] {
			cnts[i][v-'a']++
			if g[v-'a'] == nil {
				g[v-'a'] = make([]int, 0)
			}
			if len(g[v-'a']) == 0 || g[v-'a'][len(g[v-'a'])-1] != i {
				g[v-'a'] = append(g[v-'a'], i)
			}
		}
	}
	var total = 1 << m
	var dp = make([]int, total)
	for i := 1; i < len(dp); i++ {
		dp[i] = INF
	}
	for i := 0; i < total; i++ {
		if dp[i] == INF {
			continue
		}
		idx, next := 0, i
		// 找到第一个缺少的字符索引
		for j := 0; j < m; j++ {
			if (1<<j)&next == 0 {
				idx = j
				break
			}
		}
		// 枚举所有有该字符的字符串贴纸
		for _, k := range g[int(t[idx])-'a'] {
			var cnt = make([]int, len(cnts[k]))
			copy(cnt, cnts[k])
			next = i
			for j := 0; j < m; j++ {
				if (1<<j)&next != 0 {
					continue
				}
				if cnt[t[j]-'a'] > 0 {
					cnt[t[j]-'a']--
					next |= 1 << j
				}
			}
			dp[next] = min(dp[next], dp[i]+1)
		}
	}
	if dp[total-1] == INF {
		return -1
	}
	return dp[total-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
