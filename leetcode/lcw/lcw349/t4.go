package lcw349

import "sort"

/**
 * @Description https://leetcode.cn/problems/maximum-sum-queries/
 * idea: 将nums1、nums2组成的数对看作二维平面上的红点，queries数对为蓝点，问题转换为
		求蓝点(x,y)右上角值最大的红点，由于要求两个值大于等于的情况，不方便计算。所以考虑
		先将一个值【排序】，过程中求解另一个值，则能保证两者都有顺序。

		对于本题，由于较小的点实际与较大位置的点最终可能为同一个最大值，所以考虑按x坐标从大到小排序后处理，
		过程中更新y的最大值，即可保证最终求解结果正确。
		具体做法将更新与询问统一处理，最终问题变为了在>=y的所有值中寻找最大值，即区间查询，和单点更新(依次
		更新当前最大的y值)，且题目数据范围较大，使用动态开点线段树即可。

 **/

var INF = int(2e9)

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	var n = len(nums1)
	var m = len(queries)
	var points = make([][]int, 0)
	for i := 0; i < n; i++ {
		points = append(points, []int{i, 0})
	}
	for i := 0; i < m; i++ {
		points = append(points, []int{i, 1})
	}
	var getPoint = func(info []int) []int {
		if info[1] == 0 {
			return []int{nums1[info[0]], nums2[info[0]]}
		}
		return queries[info[0]]
	}
	// 先按x排序，相等再将0(更新操作)排在1(查询操作)前
	sort.Slice(points, func(i, j int) bool {
		t1, t2 := points[i], points[j]
		a, b := getPoint(t1), getPoint(t2)
		if a[0] == b[0] {
			return t1[1] < t2[1]
		}
		return a[0] > b[0]
	})
	var res = make([]int, m)
	var tree = NewSegmentTree(INF, func(lv, rv int) int {
		return max(lv, rv)
	})
	for _, info := range points {
		var p = getPoint(info)
		if info[1] == 0 {
			tree.update(p[1], p[1], p[0]+p[1])
		} else {
			var val = tree.query(p[1], INF)
			if val <= 0 {
				res[info[0]] = -1
			} else {
				res[info[0]] = val
			}
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
	return &SegmentTree{&Node{val: -1}, n, f}
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
		return 0
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
		node.left = &Node{}
	}
	if node.right == nil {
		node.right = &Node{}
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
	node.lazy = max(node.lazy, val)
	node.val = max(node.val, val)
}
