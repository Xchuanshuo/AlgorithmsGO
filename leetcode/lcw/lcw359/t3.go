package lcw359

import "sort"

func maximizeTheProfit(n int, offers [][]int) int {
	sort.Slice(offers, func(i, j int) bool {
		return offers[i][1] < offers[j][1]
	})
	var tree = NewSegmentTree(make([]int, n+1), func(lv int, rv int) int {
		return max(lv, rv)
	})
	var res = 0
	for _, o := range offers {
		var cur = max(tree.query(0, o[1]+1), tree.query(0, o[0])+o[2])
		tree.update(o[1]+1, cur)
		res = max(res, cur)
	}
	return res
}

func NewSegmentTree(data []int, fun func(lv int, rv int) int) *SegmentTree {
	tree := &SegmentTree{
		data: data,
		tree: make([]int, 4*len(data)),
		f:    fun,
	}
	tree.buildTree(0, 0, len(tree.data)-1)
	return tree
}

type SegmentTree struct {
	data []int
	tree []int
	f    func(lv, rv int) int
}

func (s *SegmentTree) leftChild(idx int) int {
	return 2*idx + 1
}

func (s *SegmentTree) rightChild(idx int) int {
	return 2*idx + 2
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

func (s *SegmentTree) query(qL, qR int) int {
	if qL < 0 || qR >= len(s.data) {
		return -1
	}
	return s._query(0, 0, len(s.data)-1, qL, qR)
}

func (s *SegmentTree) _query(treeIdx int, l, r, qL, qR int) int {
	if l == qL && r == qR {
		return s.tree[treeIdx]
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(treeIdx), s.rightChild(treeIdx)
	if qL > mid {
		return s._query(rightIdx, mid+1, r, qL, qR)
	} else if qR <= mid {
		return s._query(leftIdx, l, mid, qL, qR)
	}
	var left = s._query(leftIdx, l, mid, qL, mid)
	var right = s._query(rightIdx, mid+1, r, mid+1, qR)
	return s.f(left, right)
}

func (s *SegmentTree) update(idx int, val int) {
	if idx < 0 || idx >= len(s.data) {
		return
	}
	s.data[idx] = val
	s._update(0, 0, len(s.data)-1, idx, val)
}

func (s *SegmentTree) _update(treeIdx int, l, r, idx, val int) {
	if l == r {
		s.tree[treeIdx] = val
		return
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(treeIdx), s.rightChild(treeIdx)
	if idx > mid {
		s._update(rightIdx, mid+1, r, idx, val)
	} else {
		s._update(leftIdx, l, mid, idx, val)
	}
	s.tree[treeIdx] = s.f(s.tree[leftIdx], s.tree[rightIdx])
}
