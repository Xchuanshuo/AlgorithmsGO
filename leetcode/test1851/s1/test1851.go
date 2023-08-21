package s1

// 动态开点线段树，区间更新所有元素的最小值即可

import "math"

func minInterval(intervals [][]int, queries []int) []int {
	var tree = NewSegmentTree(int(1e7)+1, func(lv, rv int) int {
		return min(lv, rv)
	})
	for _, v := range intervals {
		tree.update(v[0], v[1], v[1]-v[0]+1)
	}
	var res = make([]int, 0)
	for _, q := range queries {
		var v = tree.query(q, q)
		if v == math.MaxInt32 {
			v = -1
		}
		res = append(res, v)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NewSegmentTree(n int, f func(lv, rv int) int) *SegmentTree {
	return &SegmentTree{&Node{}, n, f}
}

type SegmentTree struct {
	root *Node
	n    int
	f    func(lv, rv int) int
}

// Node 动态开点线段树节点信息
type Node struct {
	left, right *Node
	val         int
	lazy        int
}

func (s *SegmentTree) query(curL, curR int) int {
	return s._query(s.root, 0, s.n, curL, curR)
}

func (s *SegmentTree) update(curL, curR int, val int) {
	s._update(s.root, 0, s.n, curL, curR, val)
}

func (s *SegmentTree) _query(node *Node, l, r, curL, curR int) int {
	if r < curL || l > curR || node == nil {
		return math.MaxInt32
	}
	if l >= curL && r <= curR {
		return node.val
	}
	s.pushDown(node)
	var mid = l + (r-l)/2
	var left = s._query(node.left, l, mid, curL, curR)
	var right = s._query(node.right, mid+1, r, curL, curR)
	return s.f(left, right)
}

// 在区间[l,r]内更新[curL,curR]
func (s *SegmentTree) _update(node *Node, l, r, curL, curR int, val int) {
	if r < curL || l > curR {
		return
	}
	if l >= curL && r <= curR {
		s.do(node, val)
		return
	}
	s.pushDown(node)
	var mid = l + (r-l)/2
	s._update(node.left, l, mid, curL, curR, val)
	s._update(node.right, mid+1, r, curL, curR, val)
	node.val = s.f(node.left.val, node.right.val)
}

func (s *SegmentTree) pushDown(node *Node) {
	if node.left == nil {
		node.left = &Node{val: math.MaxInt32}
	}
	if node.right == nil {
		node.right = &Node{val: math.MaxInt32}
	}
	if node.lazy == 0 {
		return
	}
	s.do(node.left, node.lazy)
	s.do(node.right, node.lazy)
	node.lazy = 0
}

// 处理lazy逻辑
func (s *SegmentTree) do(node *Node, val int) {
	if node.lazy > 0 {
		node.lazy = min(node.lazy, val)
	} else {
		node.lazy = val
	}
	node.val = min(node.val, val)
}
