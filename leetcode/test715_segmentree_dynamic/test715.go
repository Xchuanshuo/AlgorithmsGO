package test715_segmentree_dynamic

type RangeModule struct {
	tree *SegmentTree
}

func Constructor() RangeModule {
	t := NewSegmentTree(int(1e9+1), func(lv, rv int) int {
		if lv == 1 && rv == 1 {
			return 1
		}
		return -1
	})
	t.update(0, int(1e9+1), -1)
	return RangeModule{t}
}

func (this *RangeModule) AddRange(left int, right int) {
	this.tree.update(left, right-1, 1)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	return this.tree.query(left, right-1) == 1
}

func (this *RangeModule) RemoveRange(left int, right int) {
	this.tree.update(left, right-1, -1)
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
	if r < curL || l > curR {
		return 1
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
	node.lazy = val
	node.val = val
}
