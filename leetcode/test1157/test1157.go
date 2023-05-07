package test1157

/**
 * @Description https://leetcode.cn/problems/online-majority-element-in-subarray/
 * idea: 线段树 + 摩尔投票 + 二分查找
		1.已知threshold > (r-l+1)/2, 即目标数为区间中超过半数元素
		2.使用线段树，节点记录当前值，以及在当前区间抵消投票后的数量
		3.相等的元素可能投票后【随机】返回其中一个，或左右任一区间抵消后，另一区间投票结果元素未达半数以上
 		  因此对于查找后的结果，需要二次在索引区间判定是否正确
 **/

import "sort"

type MajorityChecker struct {
	mp   map[int][]int
	tree *SegmentTree
}

func Constructor(arr []int) MajorityChecker {
	var mp = make(map[int][]int)
	for i, v := range arr {
		mp[v] = append(mp[v], i)
	}
	return MajorityChecker{mp, NewSegmentTreeByData(arr, func(lv int, rv int, cntL, cntR int) (int, int) {
		if lv == rv {
			return lv, cntL + cntR
		}
		if cntL >= cntR {
			return lv, cntL - cntR
		}
		return rv, cntR - cntL
	})}
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	v, _ := this.tree.query(left, right)
	var t = this.mp[v]
	if t == nil {
		return -1
	}
	l := sort.Search(len(t), func(j int) bool {
		return t[j] >= left
	})
	r := sort.Search(len(t), func(j int) bool {
		return t[j] > right
	})
	if r-l >= threshold {
		return v
	}
	return -1
}

func NewSegmentTree(n int, fun func(lv int, rv int, cntL, cntR int) (int, int)) *SegmentTree {
	tree := &SegmentTree{
		tree: make([]int, 4*n),
		lazy: make([]int, 4*n),
		cnt:  make([]int, 4*n),
		f:    fun,
	}
	return tree
}

func NewSegmentTreeByData(data []int, fun func(lv int, rv int, cntL, cntR int) (int, int)) *SegmentTree {
	var n = len(data)
	var tree = NewSegmentTree(n, fun)
	tree.data = data
	tree.buildTree(0, 0, n-1)
	return tree
}

func (s *SegmentTree) buildTree(treeIdx int, l, r int) {
	if l == r {
		s.tree[treeIdx] = s.data[l]
		s.cnt[treeIdx] = 1
		return
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(treeIdx), s.rightChild(treeIdx)
	s.buildTree(leftIdx, l, mid)
	s.buildTree(rightIdx, mid+1, r)
	s.tree[treeIdx], s.cnt[treeIdx] = s.f(s.tree[leftIdx], s.tree[rightIdx], s.cnt[leftIdx], s.cnt[rightIdx])
}

type SegmentTree struct {
	lazy []int
	tree []int
	cnt  []int
	data []int
	f    func(lv, rv int, cntL, cntR int) (int, int)
}

func (s *SegmentTree) leftChild(idx int) int {
	return 2*idx + 1
}

func (s *SegmentTree) rightChild(idx int) int {
	return 2*idx + 2
}

func (s *SegmentTree) query(qL, qR int) (int, int) {
	return s._query(0, 0, len(s.tree)/4-1, qL, qR)
}

func (s *SegmentTree) _query(pos int, l, r, qL, qR int) (int, int) {
	if r < qL || l > qR {
		return 0, 0
	}
	if l >= qL && r <= qR {
		return s.tree[pos], s.cnt[pos]
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(pos), s.rightChild(pos)
	left, cntL := s._query(leftIdx, l, mid, qL, qR)
	right, cntR := s._query(rightIdx, mid+1, r, qL, qR)
	return s.f(left, right, cntL, cntR)
}
