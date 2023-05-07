package test2617

import "math"

func minimumVisitedCells(g [][]int) int {
	m, n := len(g), len(g[0])
	var d1 = make([]int, m*n)
	var d2 = make([]int, m*n)
	for i := 0; i < len(d1); i++ {
		d1[i] = math.MaxInt32
	}
	copy(d2, d1)
	var tree1 = NewSegmentTreeByData(d1, func(lv int, rv int) int {
		return min(lv, rv)
	})
	var tree2 = NewSegmentTreeByData(d2, func(lv int, rv int) int {
		return min(lv, rv)
	})
	tree1.update(0, 0, 1)
	tree2.update(0, 0, 1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			s1, s2 := tree1.query(n*i+j, n*i+j), tree2.query(m*j+i, m*j+i)
			ti, tj := i+g[i][j], j+g[i][j]
			tree1.update(n*i+j+1, n*i+min(tj, n-1), min(s1, s2)+1)
			tree2.update(m*j+i+1, m*j+min(ti, m-1), min(s1, s2)+1)
		}
	}
	var res = min(tree1.query(m*n-1, m*n-1), tree2.query(m*n-1, m*n-1))
	if res == math.MaxInt32 {
		return -1
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

func NewSegmentTreeByData(data []int, fun func(lv int, rv int) int) *SegmentTree {
	var n = len(data)
	var tree = NewSegmentTree(n, fun)
	tree.data = data
	tree.buildTree(0, 0, n-1)
	return tree
}

func (s *SegmentTree) buildTree(treeIdx int, l, r int) {
	if l == r {
		s.tree[treeIdx] = s.data[l]
		return
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(treeIdx), s.rightChild(treeIdx)
	s.buildTree(leftIdx, l, mid)
	s.buildTree(rightIdx, mid+1, r)
	s.tree[treeIdx] = s.f(s.tree[leftIdx], s.tree[rightIdx])
}

type SegmentTree struct {
	lazy []int
	tree []int
	data []int
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

func (s *SegmentTree) _query(pos int, l, r, qL, qR int) int {
	if r < qL || l > qR {
		return math.MaxInt32
	}
	if l >= qL && r <= qR {
		return s.tree[pos]
	}
	s.pushDown(pos, l, r)
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(pos), s.rightChild(pos)
	var left = s._query(leftIdx, l, mid, qL, qR)
	var right = s._query(rightIdx, mid+1, r, qL, qR)
	return s.f(left, right)
}

func (s *SegmentTree) _update(pos int, l, r, curL, curR, val int) {
	if r < curL || l > curR {
		return
	}
	if l >= curL && r <= curR {
		s.do(pos, l, r, val)
		return
	}
	s.pushDown(pos, l, r)
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(pos), s.rightChild(pos)
	s._update(leftIdx, l, mid, curL, curR, val)
	s._update(rightIdx, mid+1, r, curL, curR, val)
	s.tree[pos] = s.f(s.tree[leftIdx], s.tree[rightIdx])
}

// 处理lazy逻辑
func (s *SegmentTree) do(pos int, l, r int, val int) {
	if s.lazy[pos] == 0 {
		s.lazy[pos] = val
	} else {
		s.lazy[pos] = min(s.lazy[pos], val)
	}
	s.tree[pos] = min(s.tree[pos], val)
}

func (s *SegmentTree) pushDown(pos int, l, r int) {
	if s.lazy[pos] == 0 {
		return
	}
	var mid = l + (r-l)/2
	s.do(s.leftChild(pos), l, mid, s.lazy[pos])
	s.do(s.rightChild(pos), mid+1, r, s.lazy[pos])
	s.lazy[pos] = 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
