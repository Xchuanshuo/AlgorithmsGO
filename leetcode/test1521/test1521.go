package test1521

/**
 * @Description https://leetcode.cn/problems/find-a-value-of-a-mysterious-function-closest-to-target/
 * idea: 线段树, 和->与, 对于区间的与运算， 满足新增一个数后一定不会比原来的值更加大,
		 即小区间>=大区间, 满足单调性 所以可以使用双指针，l, r维护区间左右端点
		 1.当前区间与值 > target时，需要考虑扩大区间，让与值变小，减小与target的差值, r++
		 2.当前区间值 < target时，需要减小区间，让与值变大，减小与target的差值, l++
 **/

func closestToTarget(arr []int, target int) int {
	var n = len(arr)
	var tree = NewSegmentTreeByData(arr, func(lv, rv int) int {
		return lv & rv
	})
	l, r, res := 0, 0, abs(arr[0]-target)
	for r < n {
		var cur = tree.query(l, r)
		if l > r {
			cur = -int(1e9)
		}
		res = min(res, abs(cur-target))
		if l > r || cur > target {
			r++
		} else if cur < target {
			l++
		} else {
			return 0
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
		return 0xfffffff
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
		s.do(pos, l, r)
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
func (s *SegmentTree) do(pos int, l, r int) {
	s.lazy[pos] ^= 1
	s.tree[pos] = (r - l + 1) - s.tree[pos]
}

func (s *SegmentTree) pushDown(pos int, l, r int) {
	if s.lazy[pos] == 0 {
		return
	}
	var mid = l + (r-l)/2
	s.do(s.leftChild(pos), l, mid)
	s.do(s.rightChild(pos), mid+1, r)
	s.lazy[pos] = 0
}
