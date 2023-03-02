package main

import "fmt"

type MyCalendarThree struct {
	tree *SegmentTree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		tree: NewSegmentTree(int(1e9+1), func(lv, rv int) int {
			return max(lv, rv)
		}),
	}
}

func (this *MyCalendarThree) Book(s int, e int) int {
	this.tree.update(s, e-1, 1)
	return this.tree.root.val
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
	node.lazy += val
	node.val += val
}

func main() {
	var t = Constructor()
	a := t.Book(1, 2)
	fmt.Println(a)
	b := t.Book(1, 2)
	fmt.Println(b)
}
