package test1825

/**
 * @Description https://leetcode.cn/problems/finding-mk-average/
 * idea: 由于数据范围在[1,1e5], 考虑使用两个线段树/树状数组维护区间内的个数以及区间元素和
		 1.add元素 log(N)+ log(N)时间复杂度
         2.计算平均数, 二分查找第k大的元素，最后计算前k大元素的区间和
		   答案为 (sum(m-k) - sum(k)) / (m-2*k)
 **/

type MKAverage struct {
	m, k             int
	q                []int
	cntTree, sumTree *TreeArr
}

func Constructor(m int, k int) MKAverage {
	var l = int(1e5) + 1
	return MKAverage{
		m: m, k: k, q: make([]int, 0),
		cntTree: NewTreeArr(l), sumTree: NewTreeArr(l),
	}
}

func (this *MKAverage) AddElement(num int) {
	this.q = append(this.q, num)
	this.cntTree.UpdateRange(num, num, 1)
	this.sumTree.UpdateRange(num, num, num)
	if len(this.q) > this.m {
		var f = this.q[0]
		this.q = this.q[1:]
		this.cntTree.UpdateRange(f, f, -1)
		this.sumTree.UpdateRange(f, f, -f)
	}
}

func (this *MKAverage) CalculateMKAverage() int {
	if len(this.q) < this.m {
		return -1
	}
	// 二分查找前k大的数的和
	var getSum = func(k int) int {
		l, r := 1, int(1e5)
		var mid = 0
		for l <= r {
			mid = l + (r-l)/2
			if this.cntTree.query(mid) < k {
				l = mid + 1
			} else {
				if mid == 1 || this.cntTree.query(mid-1) < k {
					break
				}
				r = mid - 1
			}
		}
		return this.sumTree.query(mid) - (this.cntTree.query(mid)-k)*mid
	}
	return (getSum(this.m-this.k) - getSum(this.k)) / (this.m - 2*this.k)
}

type TreeArr struct {
	dif1, dif2 []int
}

func NewTreeArr(n int) *TreeArr {
	return &TreeArr{dif1: make([]int, n), dif2: make([]int, n)}
}
func (t *TreeArr) QueryRange(l, r int) int {
	return t.query(r) - t.query(l-1)
}
func (t *TreeArr) UpdateRange(l, r int, val int) {
	t.update(l, val)
	t.update(r+1, -val)
}
func (t *TreeArr) query(i int) int {
	var x = i
	var res = 0
	for i > 0 {
		res += (x+1)*t.dif1[i] - t.dif2[i]
		i -= t.lowbit(i)
	}
	return res
}
func (t *TreeArr) update(i int, val int) {
	x, n := i, len(t.dif1)
	for i < n {
		t.dif1[i] += val
		t.dif2[i] += x * val
		i += t.lowbit(i)
	}
}
func (t *TreeArr) lowbit(x int) int {
	return x & -x
}

func (t *TreeArr) Clone() *TreeArr {
	var cloneDif1 = make([]int, len(t.dif1))
	var cloneDif2 = make([]int, len(t.dif2))
	copy(cloneDif1, t.dif1)
	copy(cloneDif2, t.dif2)
	return &TreeArr{cloneDif1, cloneDif2}
}
