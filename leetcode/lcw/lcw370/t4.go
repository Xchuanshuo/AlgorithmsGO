package lcw370

/**
 * @Description https://leetcode.cn/contest/weekly-contest-370/problems/maximum-balanced-subsequence-sum/
 * idea: 动态开点线段树 + dp ，移项，nums[i] - i >= nums[j] - j, 且i > j
 **/

func maxBalancedSubsequenceSum(nums []int) int64 {
	var n = len(nums)
	var offset = int(1e9) + n
	var tree = NewSegmentTree(int(1e10), func(lv, rv int64) int64 {
		return maxI64(lv, rv)
	})
	var res = int64(nums[0])
	for i, v := range nums {
		var cur = tree.query(0, v-i+offset) + int64(v)
		res = maxI64(res, cur)
		tree.update(v-i+offset, v-i+offset, cur)
	}
	return res
}

func NewSegmentTree(n int, f func(lv, rv int64) int64) *SegmentTree {
	return &SegmentTree{&Node{}, n, f}
}

type SegmentTree struct {
	root *Node
	n    int
	f    func(lv, rv int64) int64
}

// Node 动态开点线段树节点信息
type Node struct {
	left, right *Node
	val         int64
	lazy        int64
}

func (s *SegmentTree) query(curL, curR int) int64 {
	return s._query(s.root, 0, s.n, curL, curR)
}

func (s *SegmentTree) update(curL, curR int, val int64) {
	s._update(s.root, 0, s.n, curL, curR, val)
}

func (s *SegmentTree) _query(node *Node, l, r, curL, curR int) int64 {
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
func (s *SegmentTree) _update(node *Node, l, r, curL, curR int, val int64) {
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
func (s *SegmentTree) do(node *Node, val int64) {
	node.lazy += val
	node.val = maxI64(node.val, val)
}
