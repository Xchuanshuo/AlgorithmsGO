package lcwt107

import "fmt"

func minimizeConcatenatedLength(words []string) int {
	var mem = make([]map[string]int, len(words))
	for i := 0; i < len(mem); i++ {
		mem[i] = make(map[string]int)
	}
	var dfs func(i int, pre string) int
	dfs = func(i int, pre string) int {
		if i == len(words) {
			return 0
		}
		if v, exist := mem[i][pre]; exist {
			return v
		}
		var res = 0
		if pre[len(pre)-1] == words[i][0] {
			res = len(words[i]) - 1 + dfs(i+1, fmt.Sprintf("%c%c", pre[0], words[i][len(words[i])-1]))
		} else {
			res = len(words[i]) + dfs(i+1, fmt.Sprintf("%c%c", pre[0], words[i][len(words[i])-1]))
		}
		if pre[0] == words[i][len(words[i])-1] {
			res = min(res, len(words[i])-1+dfs(i+1, fmt.Sprintf("%c%c", words[i][0], pre[len(pre)-1])))
		} else {
			res = min(res, len(words[i])+dfs(i+1, fmt.Sprintf("%c%c", words[i][0], pre[len(pre)-1])))
		}
		mem[i][pre] = res
		return res
	}
	return len(words[0]) + dfs(1, words[0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
