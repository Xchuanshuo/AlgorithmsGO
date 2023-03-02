package lcwt98

/**
 * @Description https://leetcode.cn/problems/handling-sum-queries-after-update/
 * idea: 懒标记线段树
 **/

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	var n = len(nums1)
	var tree = NewSegmentTreeByData(nums1, func(lv int, rv int) int {
		return lv + rv
	})
	var sum int64 = 0
	for _, v := range nums2 {
		sum += int64(v)
	}
	var ex int64 = 0
	var res = make([]int64, 0)
	for _, q := range queries {
		var t = q[0]
		if t == 1 {
			l, r := q[1], q[2]
			tree.update(l, r, 0)
		} else if t == 2 {
			ex += int64(q[1] * tree.query(0, n-1))
		} else {
			res = append(res, sum+ex)
		}
	}
	return res
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
		return 0
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
