package test1601

func maximumRequests(n int, requests [][]int) int {
	var m = len(requests)
	var res = 0
	for i := 0; i < 1<<m; i++ {
		var ids = make([]int, 0)
		for j := 0; j < m; j++ {
			if i&(1<<j) != 0 {
				ids = append(ids, j)
			}
		}
		var degree = make([]int, n)
		for _, id := range ids {
			var req = requests[id]
			degree[req[0]]--
			degree[req[1]]++
		}
		if isValid(degree) && len(ids) > res {
			res = len(ids)
		}
	}
	return res
}

func isValid(degree []int) bool {
	for _, d := range degree {
		if d != 0 {
			return false
		}
	}
	return true
}
