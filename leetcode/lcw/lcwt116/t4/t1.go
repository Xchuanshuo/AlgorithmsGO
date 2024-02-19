package main

/**
 * @Description https://leetcode.cn/contest/biweekly-contest-116/problems/subarrays-distinct-element-sum-of-squares-ii/
 * idea: lazy线段树。维护两个值的区间修改
						1.区间节点的数值和
						2.区间节点数值的平方和 增量: (x+d)^2 - x^2 = 2dx + d^2, 其中x范围为整个区间
 **/

var M = int(1e9) + 7

func sumCounts(nums []int) int {
	var n = len(nums)
	var res = 0
	var tree = NewSegmentTree(n, func(lv, rv int) int {
		return lv + rv
	})
	var last = make(map[int]int)
	for i := 0; i < n; i++ {
		var set = make(map[int]bool)
		set[nums[i]] = true
		var p = 0 // 当前元素上一次出现的后一个位置
		if v, exist := last[nums[i]]; exist {
			p = v + 1
		}
		// [p,i]区间内任意点到i结束的子数组出现不同元素个数+1
		tree.update(p, i, 1)
		res = (res + tree.query(0, i)) % M
		last[nums[i]] = i
	}
	return res
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
	squareVal   int
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
		return node.squareVal
	}
	s.pushDown(node, l, r)
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
		s.do(node, l, r, val)
		return
	}
	s.pushDown(node, l, r)
	var mid = l + (r-l)/2
	s._update(node.left, l, mid, curL, curR, val)
	s._update(node.right, mid+1, r, curL, curR, val)
	node.val = s.f(node.left.val, node.right.val)
	node.squareVal = s.f(node.left.squareVal, node.right.squareVal)
}

func (s *SegmentTree) pushDown(node *Node, l, r int) {
	if node.left == nil {
		node.left = &Node{}
	}
	if node.right == nil {
		node.right = &Node{}
	}
	if node.lazy == 0 {
		return
	}
	var mid = l + (r-l)/2
	s.do(node.left, l, mid, node.lazy)
	s.do(node.right, mid+1, r, node.lazy)
	node.lazy = 0
}

// 处理lazy逻辑
func (s *SegmentTree) do(node *Node, l, r int, val int) {
	var cnt = r - l + 1
	node.lazy += val
	// 注意顺序, squareVal根据公式应该使用增量【之前】的节点数值和计算
	node.squareVal = (node.squareVal + 2*val*node.val + cnt*val*val%M) % M
	node.val += cnt * val
}
