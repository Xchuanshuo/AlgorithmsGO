package test_offer2_114

var graph map[byte][]byte
var inDegree []int
var isValid = true
var set map[byte]bool

func alienOrder(words []string) string {
	isValid = true
	graph = make(map[byte][]byte)
	set = make(map[byte]bool)
	inDegree = make([]int, 128)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			set[words[i][j]] = true
		}
	}
	for i := 1; i < len(words); i++ {
		addEdge(words[i-1], words[i])
		if !isValid {
			return ""
		}
	}
	var q = make([]byte, 0)
	var res []byte
	for k, _ := range set {
		if inDegree[k] == 0 {
			q = append(q, k)
			res = append(res, k)
		}
	}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, next := range graph[cur] {
			inDegree[next]--
			if inDegree[next] == 0 {
				q = append(q, next)
				res = append(res, next)
			}
		}
	}
	if len(res) == len(set) {
		return string(res)
	}
	return ""
}

func addEdge(w1, w2 string) {
	l1, l2 := len(w1), len(w2)
	var n = min(l1, l2)
	var i int
	for ; i < n; i++ {
		if w1[i] != w2[i] {
			inDegree[w2[i]]++
			graph[w1[i]] = append(graph[w1[i]], w2[i])
			return
		}
	}
	if i == n && l1 > l2 {
		isValid = false
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
