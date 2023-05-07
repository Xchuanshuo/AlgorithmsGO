package lcw340

import "math"

/**
 * @Description https://leetcode.cn/problems/minimum-number-of-visited-cells-in-a-grid/
 * idea: 线段树存储区间最小值, 单点查询 + 区间更新
	假设当前位置为(i,j) 访问格子数为step，只需把以下两个范围内的位置更新最小值为step+1

	向下移动的范围为：(i+1, j) ~ (i+grid[i][j], j)
	向右移动的范围为：(i, j+1) ~ (i, j+grid[i][j])
	注意i+grid[i][j]和j+grid[i][j]不能越界

	把二维坐标压成一维，横纵两种不同的方式压缩，用两个线段树维护
	(i,j) 横向压缩：i*n+j，纵向压缩：j*m+i
 **/

func minimumVisitedCells(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var d1 = make([]int, m*n)
	var d2 = make([]int, m*n)
	for i := 0; i < len(d1); i++ {
		d1[i] = math.MaxInt32
	}
	copy(d2, d1)
	var t1 = NewSegmentTreeByData(d1, func(lv int, rv int) int {
		return min(lv, rv)
	})
	var t2 = NewSegmentTreeByData(d2, func(lv int, rv int) int {
		return min(lv, rv)
	})
	t1.update(0, 0, 1)
	t2.update(0, 0, 1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var step = min(t1.query(i*n+j, i*n+j), t2.query(j*m+i, j*m+i))
			t1.update(i*n+j+1, i*n+min(j+grid[i][j], n-1), step+1)
			t2.update(j*m+i+1, j*m+min(i+grid[i][j], m-1), step+1)
		}
	}
	var res = min(t1.query(m*n-1, m*n-1), t2.query(m*n-1, m*n-1))
	if res == math.MaxInt32 {
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
		return math.MaxInt
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
