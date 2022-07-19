package test433

var choices = []byte{'A', 'C', 'G', 'T'}

func minMutation(start string, end string, bank []string) int {
	var set = make(map[string]bool)
	for _, v := range bank {
		set[v] = true
	}
	var res = len(start) + 1
	var dfs func(cur string, cnt int)
	var visited = make(map[string]bool)
	dfs = func(cur string, cnt int) {
		if cur == end {
			res = min(res, cnt)
			return
		}
		var bytes = []byte(cur)
		for i := 0; i < len(bytes); i++ {
			var origin = bytes[i]
			for _, v := range choices {
				if bytes[i] == v {
					continue
				}
				bytes[i] = v
				var next = string(bytes)
				if !set[next] || visited[next] {
					bytes[i] = origin
					continue
				}
				visited[next] = true
				dfs(next, cnt+1)
				visited[next] = false
				bytes[i] = origin
			}
		}
	}
	dfs(start, 0)
	if res == len(start)+1 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
