package lcw372

/**
 * @Description https://leetcode.cn/contest/weekly-contest-372/problems/find-building-where-alice-and-bob-can-meet/
 * idea: 排序 + 动态开点线段树；
		1.按顺序从大到小处理索引
		2.线段树存储每个高度的索引
		3.每个询问查询[h+1,INF]高度的最小索引
 **/

import (
	"sort"
)

var INF = int(1e9 + 1)

func leftmostBuildingQueries(hs []int, queries [][]int) []int {
	var n = len(hs)
	var as = make([][]int, 0)
	for i, q := range queries {
		if q[0] > q[1] {
			q[1], q[0] = q[0], q[1]
		}
		as = append(as, []int{q[1], i, n})
	}
	for i, v := range hs {
		as = append(as, []int{i, v, i})
	}
	sort.Slice(as, func(i, j int) bool {
		if as[i][0] != as[j][0] {
			return as[i][0] < as[j][0]
		}
		return as[i][2] < as[j][2]
	})
	var tree = NewSegmentTree(INF, func(lv, rv int) int {
		return min(lv, rv)
	})
	var res = make([]int, len(queries))
	var m = len(as)
	for i := m - 1; i >= 0; i-- {
		var a = as[i]
		if a[2] == n { // 查询
			var idx = a[1]
			var q = queries[idx]
			var h = max(hs[q[0]], hs[q[1]])
			if q[0] == q[1] || (hs[q[1]] == h && hs[q[0]] < h) {
				res[idx] = q[1]
				continue
			}
			var t = tree.query(h+1, INF)
			if t == INF {
				res[idx] = -1
			} else {
				res[idx] = t
			}
		} else { // 更新
			var h = a[1]
			tree.update(h, h, a[0])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
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
		return INF
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
		node.left = &Node{val: INF}
	}
	if node.right == nil {
		node.right = &Node{val: INF}
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
	node.lazy += val
	node.val = min(node.val, val)
}
