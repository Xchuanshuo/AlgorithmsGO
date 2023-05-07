package test1622

/**
 * @Description https://leetcode.cn/problems/fancy-sequence/
 * idea: 懒加载线段树, 乘法+加法操作 所有操作均可以写成mx+a 计算清楚乘量m和增量a的变化
		 值即可
		注意点: 乘法和加法顺序问题, 因为乘会更新加标记, 所以先乘后加，保证结果正确
        参考: https://www.luogu.com.cn/blog/user15019/solution-p3373
 **/

type Fancy struct {
	tree  *SegmentTree
	index int
}

func Constructor() Fancy {
	return Fancy{tree: NewSegmentTree(int(1e5), func(lv int, rv int) int {
		return lv + rv
	}), index: 0}
}

func (this *Fancy) Append(val int) {
	this.tree.update(this.index, this.index, val)
	this.index++
}

func (this *Fancy) AddAll(inc int) {
	this.tree.update(0, this.index-1, inc)
}

func (this *Fancy) MultAll(m int) {
	this.tree.update(0, this.index-1, -m)
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= this.index {
		return -1
	}
	return this.tree.query(idx, idx)
}

func NewSegmentTree(n int, fun func(lv int, rv int) int) *SegmentTree {
	var mulArr = make([]int, 4*n)
	for i := 0; i < len(mulArr); i++ {
		mulArr[i] = 1
	}
	tree := &SegmentTree{
		tree:    make([]int, 4*n),
		lazyAdd: make([]int, 4*n),
		lazyMul: mulArr,
		f:       fun,
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
	lazyAdd []int
	lazyMul []int
	tree    []int
	data    []int
	f       func(lv, rv int) int
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
		if val < 0 {
			s.doMul(pos, l, r, -val)
		} else {
			s.doAdd(pos, l, r, val)
		}
		return
	}
	s.pushDown(pos, l, r)
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(pos), s.rightChild(pos)
	s._update(leftIdx, l, mid, curL, curR, val)
	s._update(rightIdx, mid+1, r, curL, curR, val)
	s.tree[pos] = s.f(s.tree[leftIdx], s.tree[rightIdx])
}

// 乘lazy操作, 需要同时记录add的变化量, (mx+a)*val
// m => m*val, a => a*val, data => data*val
func (s *SegmentTree) doMul(pos int, l, r int, val int) {
	s.lazyMul[pos] = mul(s.lazyMul[pos], val)
	s.lazyAdd[pos] = mul(s.lazyAdd[pos], val)
	s.tree[pos] = mul(s.tree[pos], val)
}

// 加lazy操作 直接加即可 (mx+a)+val
// m => m, a => a+val, data => data + (r-l+1)*val
func (s *SegmentTree) doAdd(pos int, l, r int, val int) {
	s.lazyAdd[pos] += val
	s.tree[pos] = add(s.tree[pos], (r-l+1)*val)
}

func (s *SegmentTree) pushDown(pos int, l, r int) {
	if s.lazyAdd[pos] == 0 && s.lazyMul[pos] == 1 {
		return
	}
	var mid = l + (r-l)/2
	s.doMul(s.leftChild(pos), l, mid, s.lazyMul[pos])
	s.doMul(s.rightChild(pos), mid+1, r, s.lazyMul[pos])
	s.doAdd(s.leftChild(pos), l, mid, s.lazyAdd[pos])
	s.doAdd(s.rightChild(pos), mid+1, r, s.lazyAdd[pos])
	s.lazyAdd[pos] = 0
	s.lazyMul[pos] = 1
}

var M = int64(1e9) + 7

func add(a, b int) int {
	return int((int64(a) + int64(b)) % M)
}

func mul(a, b int) int {
	return int((int64(a) * int64(b)) % M)
}
