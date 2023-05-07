package test2213

/**
 * @Description https://leetcode.cn/problems/longest-substring-of-one-repeating-character/
 * idea: 线段树, 单点更新 区间查询. 用pre,suf维护区间的连续字符前后缀长度，后序遍历时进行左右区间合并(pushUp操作),
		 1.当左节点的后缀与右节点的前缀是同一字符时
		   区间的最大连续字符长度可能发生变更，同时需要维护前后缀pre,suf的值
		 2.当1不满足时，此时最大长度为左右区间最大长度取max,
		   pre为左区间前缀长度, suf为右区间后缀长度
 **/

func longestRepeating(s string, qc string, qi []int) []int {
	var data = make([]int, len(s))
	for i, v := range s {
		data[i] = int(v - 'a' + 1)
	}
	var tree = NewSegmentTreeByData(data, func(lv int, rv int) int {
		return max(lv, rv)
	})
	var res = make([]int, 0)
	for i, v := range qi {
		var tv = int(qc[i] - 'a' + 1)
		data[v] = tv
		tree.update(v, v, tv)
		res = append(res, tree.tree[0])
	}
	return res
}

func NewSegmentTree(n int, fun func(lv int, rv int) int) *SegmentTree {
	tree := &SegmentTree{
		tree: make([]int, 4*n),
		pre:  make([]int, 4*n),
		suf:  make([]int, 4*n),
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
		s.tree[treeIdx], s.pre[treeIdx], s.suf[treeIdx] = 1, 1, 1
		return
	}
	var mid = l + (r-l)/2
	s.buildTree(s.leftChild(treeIdx), l, mid)
	s.buildTree(s.rightChild(treeIdx), mid+1, r)
	s.pushUp(treeIdx, l, r)
}

func (s *SegmentTree) pushUp(treeIdx int, l, r int) {
	leftIdx, rightIdx := s.leftChild(treeIdx), s.rightChild(treeIdx)
	var mid = l + (r-l)/2
	s.tree[treeIdx] = s.f(s.tree[leftIdx], s.tree[rightIdx])
	s.pre[treeIdx], s.suf[treeIdx] = s.pre[leftIdx], s.suf[rightIdx]
	if s.data[mid] == s.data[mid+1] {
		if s.pre[leftIdx] == mid-l+1 { // 左节点前缀+右节点前缀
			s.pre[treeIdx] = s.pre[leftIdx] + s.pre[rightIdx]
		}
		if s.suf[rightIdx] == r-mid { // 左节点后缀+右节点后缀
			s.suf[treeIdx] = s.suf[leftIdx] + s.suf[rightIdx]
		}
		s.tree[treeIdx] = max(s.tree[treeIdx], s.suf[leftIdx]+s.pre[rightIdx])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type SegmentTree struct {
	tree []int
	pre  []int
	suf  []int
	data []int
	f    func(lv, rv int) int
}

func (s *SegmentTree) leftChild(idx int) int {
	return 2*idx + 1
}

func (s *SegmentTree) rightChild(idx int) int {
	return 2*idx + 2
}

func (s *SegmentTree) update(qL, qR int, val int) {
	s._update(0, 0, len(s.tree)/4-1, qL, qR, val)
}

func (s *SegmentTree) _update(pos int, l, r, curL, curR, val int) {
	if r < curL || l > curR {
		return
	}
	if l >= curL && r <= curR {
		return
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(pos), s.rightChild(pos)
	s._update(leftIdx, l, mid, curL, curR, val)
	s._update(rightIdx, mid+1, r, curL, curR, val)
	s.pushUp(pos, l, r)
}
