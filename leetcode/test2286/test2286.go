package test2286

/**
 * @Description https://leetcode.cn/problems/booking-concert-tickets-in-groups/
 * idea: 线段树 + 二分, 使用线段树维护n排座位，因为两种排座方式只要有位置 k位观众尽量
		连续坐在一起 所以需要维护每排已经坐了的观众数量
			1.支持任意区间的观众总数查询sum
			2.支持任意排(区间)已坐最少观众数的查询min
		维护好sum和min后，对于Gather操作，需要在线段树上进行二分查找，q次Gather整体复杂度为O(q*log(n))
			a.对于区间[0,maxRow] 如果min > m-k, 说明前maxRow排没有座位可以坐下k位观众 返回空
			b.如果min <= m-k, 优先左半边区间查找, 剩余座位不足再去右边查找
			c.查找到最左边的区间即为最靠前的排，已坐观众数+1，即为k位观众的开始位置
		对于Scatter操作，由于需要从[0,maxRow]中未满座开始坐起, 所以二分找到第一个有剩余位置的排数,
		依次填充即可，整体时间复杂度 O((n+q)*log(n))
		若[0,maxRow]的sum>(maxRow-1)*m+m-k 说明剩余座位不足k个，返回false即可
 **/

type BookMyShow struct {
	n, m int
	tree *SegmentTree
}

func Constructor(n int, m int) BookMyShow {
	return BookMyShow{
		n, m, NewSegmentTree(n),
	}
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	if k > this.m {
		return []int{}
	}
	var idx = this.tree.search(0, maxRow, this.m-k)
	if idx == -1 {
		return []int{}
	}
	var seats = this.tree.query(idx, idx)
	this.tree.update(idx, idx, k) // 安排连坐k个人
	return []int{idx, seats}
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	if (maxRow+1)*this.m-this.tree.query(0, maxRow) < k {
		return false
	}
	var idx = this.tree.search(0, maxRow, this.m-1)
	for idx <= maxRow {
		var cur = this.m - this.tree.query(idx, idx)
		if cur >= k {
			this.tree.update(idx, idx, k)
			break
		}
		k -= cur
		this.tree.update(idx, idx, cur)
		idx = this.tree.search(idx+1, maxRow, this.m-1)
	}
	return true
}

func NewSegmentTree(n int) *SegmentTree {
	tree := &SegmentTree{
		sum: make([]int, 4*n),
		min: make([]int, 4*n),
	}
	return tree
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type SegmentTree struct {
	min []int
	sum []int
	m   int
	f   func(lv, rv int) int
}

func (s *SegmentTree) leftChild(idx int) int {
	return 2*idx + 1
}

func (s *SegmentTree) rightChild(idx int) int {
	return 2*idx + 2
}

func (s *SegmentTree) query(qL, qR int) int {
	return s._query(0, 0, len(s.sum)/4-1, qL, qR)
}

func (s *SegmentTree) update(qL, qR int, val int) {
	s._update(0, 0, len(s.sum)/4-1, qL, qR, val)
}

func (s *SegmentTree) _query(pos int, l, r, qL, qR int) int {
	if r < qL || l > qR {
		return 0
	}
	if l >= qL && r <= qR {
		return s.sum[pos]
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(pos), s.rightChild(pos)
	var left = s._query(leftIdx, l, mid, qL, qR)
	var right = s._query(rightIdx, mid+1, r, qL, qR)
	return left + right
}

func (s *SegmentTree) _update(pos int, l, r, curL, curR, val int) {
	if r < curL || l > curR {
		return
	}
	if l == r {
		s.sum[pos] += val
		s.min[pos] += val
		return
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(pos), s.rightChild(pos)
	s._update(leftIdx, l, mid, curL, curR, val)
	s._update(rightIdx, mid+1, r, curL, curR, val)
	s.sum[pos] = s.sum[leftIdx] + s.sum[rightIdx]
	s.min[pos] = min(s.min[leftIdx], s.min[rightIdx])
}

// 在[l,r]排座位中搜索第一个剩余k个位置的排数
func (s *SegmentTree) search(qL, qR int, k int) int {
	return s._search(0, 0, len(s.sum)/4-1, qL, qR, k)
}

func (s *SegmentTree) _search(pos int, l, r int, curL, curR, val int) int {
	if s.min[pos] > val {
		return -1
	}
	if l == r {
		return l
	}
	var mid = l + (r-l)/2
	leftIdx, rightIdx := s.leftChild(pos), s.rightChild(pos)
	if s.min[leftIdx] <= val {
		return s._search(leftIdx, l, mid, curL, curR, val)
	}
	if s.min[rightIdx] <= val && mid < curR {
		return s._search(rightIdx, mid+1, r, curL, curR, val)
	}
	return -1
}
