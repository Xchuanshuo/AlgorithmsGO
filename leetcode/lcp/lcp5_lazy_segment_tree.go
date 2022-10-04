package main

const M = int(1e9) + 7

func bonus(n int, leadership [][]int, operations [][]int) []int {
	var g = make(map[int][]int)
	for _, ld := range leadership {
		g[ld[0]] = append(g[ld[0]], ld[1])
	}
	L, R := make(map[int]int), make(map[int]int)
	var pos = -1
	var dfs func(int)
	dfs = func(id int) {
		pos++
		L[id] = pos
		for _, s := range g[id] {
			dfs(s)
		}
		R[id] = pos
	}
	dfs(1)
	tree := NewSegmentTree(n, func(lv int, rv int) int {
		return (lv + rv) % M
	})
	var res = make([]int, 0)
	for _, op := range operations {
		if op[0] == 1 {
			id, cnt := op[1], op[2]
			tree.update(L[id], L[id], cnt)
		} else if op[0] == 2 {
			id, cnt := op[1], op[2]
			tree.update(L[id], R[id], cnt)
		} else {
			res = append(res, tree.query(L[op[1]], R[op[1]]))
		}
	}
	return res
}

func NewSegmentTree(n int, fun func(lv int, rv int) int) *SegmentTree {
	tree := &SegmentTree{
		tree: make([]int, 4*n),
		lazy: make([]int, 4*n),
		f:    fun,
	}
	return tree
}

type SegmentTree struct {
	lazy []int
	tree []int
	f    func(lv, rv int) int
}

func (s *SegmentTree) leftChild(idx int) int {
	return 2*idx + 1
}

func (s *SegmentTree) rightChild(idx int) int {
	return 2*idx + 2
}

func (s *SegmentTree) query(qL, qR int) int {
	return s._query(0, 0, len(s.tree)/4-1, qL, qR)
}

func (s *SegmentTree) update(qL, qR int, val int) {
	s._update(0, 0, len(s.tree)/4-1, qL, qR, val)
}

func (s *SegmentTree) pushDown(pos, l, r int) {
	if s.lazy[pos] == 0 {
		return
	}
	// 将lazy转换为当前节点的值
	s.tree[pos] += s.lazy[pos] * (r - l + 1)
	if l < r { // 传递lazy到下一层
		s.lazy[s.leftChild(pos)] += s.lazy[pos]
		s.lazy[s.rightChild(pos)] += s.lazy[pos]
	}
	s.lazy[pos] = 0
}

func (s *SegmentTree) _query(pos int, l, r, qL, qR int) int {
	if r < qL || l > qR {
		return 0
	}
	s.pushDown(pos, l, r)
	if l >= qL && r <= qR {
		return s.tree[pos]
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(pos), s.rightChild(pos)
	var left = s._query(leftIdx, l, mid, qL, qR)
	var right = s._query(rightIdx, mid+1, r, qL, qR)
	return s.f(left, right)
}

func (s *SegmentTree) _update(pos int, l, r, curL, curR, val int) {
	s.pushDown(pos, l, r)
	if r < curL || l > curR {
		return
	}
	if l >= curL && r <= curR {
		s.lazy[pos] = val
		s.pushDown(pos, l, r)
		return
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(pos), s.rightChild(pos)
	s._update(leftIdx, l, mid, curL, curR, val)
	s._update(rightIdx, mid+1, r, curL, curR, val)
	s.tree[pos] = s.f(s.tree[leftIdx], s.tree[rightIdx])
}
