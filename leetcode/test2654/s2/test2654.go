package s1

/**
 * @Description https://leetcode.cn/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1/
 * idea: 解法1关键为枚举区间，且求数组gcd满足单调性，所以考虑使用二分查找长度最小的区间。
		 对于每个区间的gcd，由于满足结合律，可以使用线段树进行区间gcd的计算。
		 整体时间复杂度为O(n*log(N^2)) 其中两个log分别为二分区间枚举和区间gcd计算
 **/
func minOperations(nums []int) int {
	var n = len(nums)
	var cnt = 0
	for _, v := range nums {
		if v == 1 {
			cnt++
		}
	}
	if cnt > 0 {
		return n - cnt
	}
	var tree = NewSegmentTree(nums, func(lv int, rv int) int {
		return gcd(lv, rv)
	})
	var check = func(l int) bool {
		for i := 0; i <= n-l; i++ {
			var j = i + l - 1
			if tree.query(i, j) == 1 {
				return true
			}
		}
		return false
	}
	l, r := 1, n
	for l <= r {
		var mid = l + (r-l)/2
		if !check(mid) {
			l = mid + 1
		} else {
			if mid == 1 || !check(mid-1) {
				return n - 2 + mid
			}
			r = mid - 1
		}
	}
	return -1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func NewSegmentTree(data []int, fun func(lv int, rv int) int) *SegmentTree {
	tree := &SegmentTree{
		data: data,
		tree: make([]int, 4*len(data)),
		f:    fun,
	}
	tree.buildTree(0, 0, len(tree.data)-1)
	return tree
}

type SegmentTree struct {
	data []int
	tree []int
	f    func(lv, rv int) int
}

func (s *SegmentTree) leftChild(idx int) int {
	return 2*idx + 1
}

func (s *SegmentTree) rightChild(idx int) int {
	return 2*idx + 2
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

func (s *SegmentTree) query(qL, qR int) int {
	if qL < 0 || qR >= len(s.data) {
		return -1
	}
	return s._query(0, 0, len(s.data)-1, qL, qR)
}

func (s *SegmentTree) _query(treeIdx int, l, r, qL, qR int) int {
	if l == qL && r == qR {
		return s.tree[treeIdx]
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(treeIdx), s.rightChild(treeIdx)
	if qL > mid {
		return s._query(rightIdx, mid+1, r, qL, qR)
	} else if qR <= mid {
		return s._query(leftIdx, l, mid, qL, qR)
	}
	var left = s._query(leftIdx, l, mid, qL, mid)
	var right = s._query(rightIdx, mid+1, r, mid+1, qR)
	return s.f(left, right)
}

func (s *SegmentTree) update(idx int, val int) {
	if idx < 0 || idx >= len(s.data) {
		return
	}
	s.data[idx] = val
	s._update(0, 0, len(s.data)-1, idx, val)
}

func (s *SegmentTree) _update(treeIdx int, l, r, idx, val int) {
	if l == r {
		s.tree[treeIdx] = val
		return
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(treeIdx), s.rightChild(treeIdx)
	if idx > mid {
		s._update(rightIdx, mid+1, r, idx, val)
	} else {
		s._update(leftIdx, l, mid, idx, val)
	}
	s.tree[treeIdx] = s.f(s.tree[leftIdx], s.tree[rightIdx])
}
