package test854_astar

import (
	"container/heap"
	"sort"
)

var dis = make(map[string]int)
var target string

func kSimilarity(s1 string, s2 string) int {
	dis = make(map[string]int)
	target = s2
	h := &Heap{}
	dis[s1] = 0
	heap.Push(h, s1)
	for h.Len() > 0 {
		var cur = heap.Pop(h).(string)
		var step = dis[cur]
		if cur == s2 {
			return step
		}
		for _, nxt := range next(cur, s2) {
			if _, exist := dis[nxt]; !exist || dis[nxt] > step+1 {
				dis[nxt] = step + 1
				heap.Push(h, nxt)
			}
		}
	}
	return 0
}

func next(s1, s2 string) []string {
	var i = 0
	for i < len(s1) && s1[i] == s2[i] {
		i++
	}
	var next = make([]string, 0)
	var bytes = []byte(s1)
	for j := i + 1; j < len(s1); j++ {
		if s1[j] == s2[i] && s1[j] != s2[j] {
			bytes[i], bytes[j] = bytes[j], bytes[i]
			next = append(next, string(bytes))
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}
	}
	return next
}

func g(a, b string) int {
	var cnt = 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
		}
	}
	return (cnt + 1) / 2
}

type Heap struct {
	sort.StringSlice
}

func (h Heap) Less(i, j int) bool {
	a, b := h.StringSlice[i], h.StringSlice[j]
	va, vb := g(a, target), g(b, target)
	disA, disB := dis[a], dis[b]
	return (va + disA) < (vb + disB)
}

func (h *Heap) Push(x interface{}) {
	h.StringSlice = append(h.StringSlice, x.(string))
}

func (h *Heap) Pop() interface{} {
	a := h.StringSlice
	x := a[len(a)-1]
	h.StringSlice = a[0 : len(a)-1]
	return x
}

//func kSimilarity(s1 string, s2 string) int {
//	var level = 0
//	var q = make([]string, 0)
//	q = append(q, s1)
//	var visited = make(map[string]bool)
//	visited[s1] = true
//	for len(q) > 0 {
//		var sz = len(q)
//		for i := 0; i < sz; i++ {
//			var cur = q[0]
//			q = q[1:]
//			if s2 == cur {
//				return level
//			}
//			for _, nxt := range next(cur, s2) {
//				if visited[nxt] {
//					continue
//				}
//				q = append(q, nxt)
//				visited[nxt] = true
//			}
//		}
//		level++
//	}
//	return level
//}
